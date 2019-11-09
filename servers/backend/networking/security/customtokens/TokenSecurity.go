package customtokens

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

//GenerateSignature returns a HS256 signature used for tokenization
func GenerateSignature(data, secret []byte) (string, error) {
	h := hmac.New(sha256.New, secret)
	_, err := h.Write(data)
	if err != io.ErrClosedPipe {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
