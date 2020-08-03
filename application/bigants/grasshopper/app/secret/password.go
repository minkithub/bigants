package secret

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// https://github.com/django/django/blob/master/django/contrib/auth/hashers.py
const algorithm = "argon2i"
const timeCost uint32 = 2
const memoryCost uint32 = 512
const parallelism uint8 = 2
const saltLength uint32 = 16
const hashLength uint32 = 16

type encodedPassword struct {
	memoryCost  uint32
	timeCost    uint32
	parallelism uint8
	salt        []byte
	hash        []byte
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func CreateHashFromPassword(password string) (encodedHash string, err error) {
	salt, err := generateRandomBytes(saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.Key([]byte(password), salt, timeCost, memoryCost, parallelism, hashLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf(
		"argon2$%s$v=%d$m=%d,t=%d,p=%d$%s$%s",
		algorithm,
		argon2.Version,
		memoryCost,
		timeCost,
		parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	p, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.Key(
		[]byte(password),
		p.salt,
		p.timeCost,
		p.memoryCost,
		p.parallelism,
		hashLength,
	)

	if subtle.ConstantTimeCompare(p.hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (*encodedPassword, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, ErrInvalidHash
	}
	var p encodedPassword

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, err
	}
	if version != argon2.Version {
		return nil, ErrIncompatibleVersion
	}

	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memoryCost, &p.timeCost, &p.parallelism)
	if err != nil {
		return nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, err
	}
	p.salt = salt

	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, err
	}
	p.hash = hash

	return &p, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
