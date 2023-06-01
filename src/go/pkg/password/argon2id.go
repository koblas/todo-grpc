package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var _ Hasher = Argon2ID{}
var _ Verifier = Argon2ID{}

const ARGON_NAME = "argon2id"

type Argon2ID struct {
	format  string
	version int
	time    uint32
	memory  uint32
	keyLen  uint32
	saltLen uint32
	threads uint8
}

func NewArgon2ID() Argon2ID {
	// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
	// From the Go docs --
	//    The draft RFC recommends[2] time=3, and memory=32*1024 is a sensible number.
	//    If using that amount of memory (32 MB) is not possible in some contexts then the
	//    time parameter can be increased to compensate.

	return Argon2ID{
		format:  "$" + ARGON_NAME + "$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version: argon2.Version,
		time:    3,
		memory:  32 * 1024,
		keyLen:  32,
		saltLen: 16,
		threads: 4,
	}
}

func (a Argon2ID) Hash(plain string) (string, error) {
	salt := make([]byte, a.saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := a.hasher(plain, salt, a.keyLen)

	return fmt.Sprintf(
			a.format,
			a.version,
			a.memory,
			a.time,
			a.threads,
			base64.RawStdEncoding.EncodeToString(salt),
			base64.RawStdEncoding.EncodeToString(hash),
		),
		nil
}

func (a Argon2ID) Verify(plain, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")

	if len(hashParts) != 6 || hashParts[0] != "" {
		return false, errors.New("invalid hashed password")
	}
	if hashParts[1] != ARGON_NAME {
		return false, errors.New("unexpected hasher")
	}

	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &a.memory, &a.time, &a.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false, err
	}

	// we use the original's keylength to protected against "upgrades" in the comparison
	hashToCompare := a.hasher(plain, salt, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

func (a Argon2ID) hasher(plain string, salt []byte, keyLen uint32) []byte {
	return argon2.IDKey([]byte(plain), salt, a.time, a.memory, a.threads, keyLen)
}
