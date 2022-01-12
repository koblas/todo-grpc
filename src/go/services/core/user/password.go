package user

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"

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

/** This is not a secure token value */
func hmacCreate(key, token string) ([]byte, string, error) {
	mac := hmac.New(sha256.New, []byte(key))

	_, err := mac.Write([]byte(token))
	if err != nil {
		return nil, token, err
	}

	return mac.Sum(nil), token, nil
}

/** Compare non-secure token values */
func hmacCompare(key, token string, check []byte) (bool, error) {
	value, _, err := hmacCreate(key, token)
	if err != nil {
		return false, err
	}

	return bytes.Equal(value, check), nil
}
