package crypto

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type ArgonHasher struct{}

func NewArgonHasher() *ArgonHasher {
	return &ArgonHasher{}
}

func (a *ArgonHasher) Hash(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(append(salt, hash...)), nil
}

func (a *ArgonHasher) Verify(password, encoded string) bool {
	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil || len(data) < 16 {
		return false
	}

	salt := data[:16]
	expected := data[16:]
	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 4, 32)

	if len(hash) != len(expected) {
		return false
	}

	var result byte
	for i := 0; i < len(hash); i++ {
		result |= hash[i] ^ expected[i]
	}
	return result == 0
}
