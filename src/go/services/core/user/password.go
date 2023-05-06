package user

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(password string) string {
	if len(password) < 8 {
		return "length_too_short"
	}

	return ""
}

func passwordEncrypt(password string) ([]byte, string) {
	msg := validatePassword(password)
	if msg != "" {
		return nil, msg
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "encryption_failed"
	}

	return pass, ""
}

func tokenEncrypt(password string) ([]byte, string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return enc, password, err
}

func passwordCompare(hashedPassword []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}

// Generate a random string of a given length
const LETTERS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

var LETTERS_SIZE = big.NewInt(int64(len(LETTERS)))

func randomBytes(n int) ([]byte, error) {
	maxValue := big.NewInt(int64(256))

	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, maxValue)
		if err != nil {
			return nil, err
		}
		ret[i] = byte(num.Int64())
	}

	return ret, nil

}

func randomString(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, LETTERS_SIZE)
		if err != nil {
			return "", err
		}
		ret[i] = LETTERS[num.Int64()]
	}

	return string(ret), nil

}

/** This is not a secure token value */
func hmacCreate(key []byte, token string) ([]byte, error) {
	mac := hmac.New(sha256.New, []byte(key))

	_, err := mac.Write([]byte(token))
	if err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

/** Compare non-secure token values */
func hmacCompare(key []byte, token string, check []byte) (bool, error) {
	value, err := hmacCreate(key, token)
	if err != nil {
		return false, err
	}

	return bytes.Equal(value, check), nil
}
