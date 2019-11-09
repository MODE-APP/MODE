package customtokens

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

//Token represents the parts inside of the custom tokens
type Token struct {
	TokenHeader    map[string]string
	TokenPayload   map[string]string
	TokenSignature string
}

//BasicToken is used for testing purposes before the mapping of things inside of the token are done
type BasicToken struct {
	TokenHeader    string
	TokenPayload   string
	TokenSignature string
}

//GenerateSignature returns a HS256 signature used for tokenization
func GenerateSignature(data, secret []byte) (string, error) {
	h := hmac.New(sha256.New, secret)
	_, err := h.Write(data)
	if err != io.ErrClosedPipe && err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
