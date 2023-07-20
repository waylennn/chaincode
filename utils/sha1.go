package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// ComputeSHA1 calculate the sha1 value of binary data
func ComputeSHA1(data []byte) ([]byte, error) {
	h := sha1.New()
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}

	bs := h.Sum(nil)
	return bs, nil
}

// ComputeSHA1String calculate the sha1 value of binary data
func ComputeSHA1String(data []byte) (string, error) {
	sha1, err := ComputeSHA1(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha1), nil
}
