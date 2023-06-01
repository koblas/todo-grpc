package password

import (
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var _ Hasher = Bcrypt{}
var _ Verifier = Bcrypt{}

const BCRYPT_NAME = "bcrypt"

type Bcrypt struct {
	cost int
}

func NewBcrypt() Bcrypt {
	return Bcrypt{
		cost: bcrypt.DefaultCost,
	}
}

func (b Bcrypt) Hash(password string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)

	return "$" + BCRYPT_NAME + "$" + base64.RawStdEncoding.EncodeToString(enc), err
}

func (Bcrypt) Verify(plain, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	if len(hashParts) != 3 || hashParts[0] != "" {
		return false, errors.New("invalid hashed password")
	}
	if hashParts[1] != BCRYPT_NAME {
		return false, errors.New("unexpected hasher")
	}

	hashedPassword, err := base64.RawStdEncoding.DecodeString(hashParts[2])
	if err != nil {
		return false, err
	}

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(plain)) == nil, nil
}
