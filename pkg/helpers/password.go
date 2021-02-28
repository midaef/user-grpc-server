package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// PasswordHasher ...
type PasswordHasher interface {
	NewMD5Hash() (string, error)
}

// Password ...
type Md5 struct {
	Salt string
}

// NewHash ...
func NewHasher(salt string) *Md5 {
	return &Md5{
		Salt: salt,
	}
}

// NewMD5Hash ...
func (m *Md5) NewMD5Hash(pass string) (string, error) {
	hasher := md5.New()

	_, err := hasher.Write([]byte(pass))
	if err != nil {
		return "", errors.New("pass-hasher: hasher error")
	}

	return hex.EncodeToString(append(hasher.Sum(nil), []byte(m.Salt)...)), nil
}
