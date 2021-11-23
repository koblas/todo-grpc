package key_manager

import (
	cryptoaes "crypto/aes"
	"encoding/base64"
	"io"
	"log"
	"regexp"

	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

type environmentKms struct {
	Encoder
	Decoder

	key []byte
}

// 32 bytes is 256 bits for AES256
const envKeySize = 32

type encryptedValue struct {
	data     []byte
	iv       []byte
	tag      []byte
	datatype string
}

// Borrowed from sops
var encre = regexp.MustCompile(`^ENC\[AES256_GCM,data:(.+),iv:(.+),tag:(.+),type:(.+)\]`)

func NewSecureEnvironment() KeyManager {
	secret := []byte(os.Getenv("KMS_SECRET"))

	// Technically this should fail, but ....
	if len(secret) == 0 {
		secret = make([]byte, envKeySize)
		_, err := rand.Read(secret)
		if err != nil {
			log.Fatal("Unable to build default secret")
		}
	}

	return &environmentKms{
		key: secret,
	}
}

// Return the
//   -- encrypted key
//   -- plaintext
func (kms *environmentKms) createDataKey() ([]byte, []byte, error) {
	aescipher, err := cryptoaes.NewCipher(kms.key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(aescipher)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create GCM: %s", err)
	}
	nonceSize := gcm.NonceSize()

	nonce := make([]byte, nonceSize)
	value := make([]byte, envKeySize)

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	if _, err = io.ReadFull(rand.Reader, value); err != nil {
		panic(err.Error())
	}

	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to
	//  the encrypted data. The first nonce argument in Seal is the prefix.
	out := gcm.Seal(nonce, nonce, value, nil)

	return out, value, nil
}

func (kms *environmentKms) decodeDataKey(datakey []byte) ([]byte, error) {
	aescipher, err := cryptoaes.NewCipher(kms.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aescipher)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()

	nonce, data := datakey[:nonceSize], datakey[nonceSize:]
	decryptedBytes, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt with AES_GCM: %s", err)
	}

	return decryptedBytes, nil
}

func parse(value string) (*encryptedValue, error) {
	matches := encre.FindStringSubmatch(value)
	if matches == nil {
		return nil, fmt.Errorf("input string %s does not match data format", value)
	}
	data, err := base64.StdEncoding.DecodeString(matches[1])
	if err != nil {
		return nil, fmt.Errorf("error base64-decoding data: %s", err)
	}
	iv, err := base64.StdEncoding.DecodeString(matches[2])
	if err != nil {
		return nil, fmt.Errorf("error base64-decoding iv: %s", err)
	}
	tag, err := base64.StdEncoding.DecodeString(matches[3])
	if err != nil {
		return nil, fmt.Errorf("error base64-decoding tag: %s", err)
	}
	datatype := string(matches[4])

	return &encryptedValue{data, iv, tag, datatype}, nil
}

func (skms *environmentKms) Encode(value []byte) (SecureValue, error) {
	if len(value) == 0 {
		return SecureValue{
			KmsUri:  "localenv:KMS_SECRET",
			DataKey: []byte{},
			Data:    value,
		}, nil
	}

	dataKeyEnc, dataKeyPlain, err := skms.createDataKey()
	if err != nil {
		return SecureValue{}, err
	}

	additionalData := []byte{}

	aescipher, err := cryptoaes.NewCipher(dataKeyPlain)
	if err != nil {
		return SecureValue{}, fmt.Errorf("could not initialize AES GCM encryption cipher: %s", err)
	}

	gcm, err := cipher.NewGCM(aescipher)
	if err != nil {
		return SecureValue{}, fmt.Errorf("could not create GCM: %s", err)
	}
	nonceSize := gcm.NonceSize()
	iv := make([]byte, nonceSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}

	out := gcm.Seal(nil, iv, value, additionalData)

	strValue := fmt.Sprintf("ENC[AES256_GCM,data:%s,iv:%s,tag:%s,type:%s]",
		base64.StdEncoding.EncodeToString(out[:len(out)-cryptoaes.BlockSize]),
		base64.StdEncoding.EncodeToString(iv),
		base64.StdEncoding.EncodeToString(out[len(out)-cryptoaes.BlockSize:]),
		"byte")

	return SecureValue{
		KmsUri:  "localenv:KMS_SECRET",
		DataKey: dataKeyEnc,
		Data:    []byte(strValue),
	}, nil
}

func (smks *environmentKms) Decode(value SecureValue) ([]byte, error) {
	if len(value.DataKey) == 0 || len(value.Data) == 0 {
		return value.Data, nil
	}

	encryptedValue, err := parse(string(value.Data))
	if err != nil {
		return nil, err
	}

	dataPlain, err := smks.decodeDataKey(value.DataKey)
	if err != nil {
		return nil, err
	}
	aescipher, err := cryptoaes.NewCipher(dataPlain)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(aescipher, len(encryptedValue.iv))
	if err != nil {
		return nil, err
	}
	additionalData := []byte{}
	data := append(encryptedValue.data, encryptedValue.tag...)
	decryptedBytes, err := gcm.Open(nil, encryptedValue.iv, data, []byte(additionalData))
	if err != nil {
		return nil, fmt.Errorf("could not decrypt with AES_GCM: %s", err)
	}

	if encryptedValue.datatype != "byte" {
		return nil, fmt.Errorf("should only be string")
	}

	return decryptedBytes, err
}
