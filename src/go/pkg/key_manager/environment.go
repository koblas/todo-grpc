package key_manager

import (
	cryptoaes "crypto/aes"
	"crypto/sha1"
	"errors"
	"strings"

	"encoding/base64"
	"io"
	"log"

	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

type environmentKms struct {
	Encoder
	Decoder

	key    []byte
	keyUri string
}

// 32 bytes is 256 bits for AES256
const envKeySize = 32

var (
	ErrorKeyLookupFailed = errors.New("kms uri key not found")
	ErrorDecodeParse     = errors.New("unable to parse data")
)

type encryptedValue struct {
	data     []byte
	iv       []byte
	tag      []byte
	datatype string
}

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

	hasher := sha1.New()
	hasher.Write(secret)

	return &environmentKms{
		keyUri: "localenv:" + base64.URLEncoding.EncodeToString(hasher.Sum(nil)),
		key:    secret,
	}
}

// Return the
//
//	-- encrypted key
//	-- plaintext
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
	if !strings.HasPrefix(value, "ENC[") || !strings.HasSuffix(value, "]") {
		return nil, fmt.Errorf("%w: input string %s does not match data format", ErrorDecodeParse, value)
	}
	value = value[4 : len(value)-1]
	parts := strings.Split(value, ",")

	if parts[0] != "AES256_GCM" {
		return nil, fmt.Errorf("%w: unexpected algorithm %s", ErrorDecodeParse, parts[0])
	}

	eValue := encryptedValue{}
	for _, part := range parts[1:] {
		subParts := strings.SplitN(part, ":", 2)
		if len(subParts) != 2 {
			return nil, fmt.Errorf("%w: expected key:value %s", ErrorDecodeParse, part)
		}
		var err error
		switch subParts[0] {
		case "data":
			eValue.data, err = base64.StdEncoding.DecodeString(subParts[1])
			if err != nil {
				return nil, fmt.Errorf("%w: error base64-decoding data: %s", ErrorDecodeParse, err)
			}
		case "iv":
			eValue.iv, err = base64.StdEncoding.DecodeString(subParts[1])
			if err != nil {
				return nil, fmt.Errorf("%w: error base64-decoding data: %s", ErrorDecodeParse, err)
			}
		case "tag":
			eValue.tag, err = base64.StdEncoding.DecodeString(subParts[1])
			if err != nil {
				return nil, fmt.Errorf("%w: error base64-decoding data: %s", ErrorDecodeParse, err)
			}
		case "type":
			eValue.datatype = subParts[1]
		}
	}

	return &eValue, nil
}

func (skms *environmentKms) Encode(value []byte) (SecureValue, error) {
	dataKeyEnc, dataKeyPlain, err := skms.createDataKey()
	if err != nil {
		return SecureValue{}, err
	}

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

	out := gcm.Seal(nil, iv, value, nil)

	strValue := fmt.Sprintf("ENC[AES256_GCM,data:%s,iv:%s,tag:%s,type:%s]",
		base64.StdEncoding.EncodeToString(out[:len(out)-cryptoaes.BlockSize]),
		base64.StdEncoding.EncodeToString(iv),
		base64.StdEncoding.EncodeToString(out[len(out)-cryptoaes.BlockSize:]),
		"byte")

	return SecureValue{
		KmsUri:  skms.keyUri,
		DataKey: dataKeyEnc,
		Data:    []byte(strValue),
	}, nil
}

func (smks *environmentKms) Decode(value SecureValue) ([]byte, error) {
	if value.KmsUri != smks.keyUri {
		return nil, ErrorKeyLookupFailed
	}

	if len(value.Data) == 0 {
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
		return nil, fmt.Errorf("should only be byte type")
	}

	if decryptedBytes == nil {
		return []byte{}, err
	}

	return decryptedBytes, err
}
