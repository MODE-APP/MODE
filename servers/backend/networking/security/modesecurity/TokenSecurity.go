package modesecurity

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

//Token represents the parts inside of the custom tokens
type Token struct {
	TokenHeader    map[string]string
	TokenPayload   map[string]string
	TokenSignature string
}

//BasicToken is used for testing purposes before the mapping of things inside of the token are done

//GenerateSignature returns a HS256 signature used for tokenization
func generateHS256(data, secret []byte) (string, error) {
	h := hmac.New(sha256.New, secret)
	_, err := h.Write(data)
	if err != io.ErrClosedPipe && err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

//ValidateToken validates the token's signature
func ValidateToken(token *protos.SignedToken, key *os.File) error {
	sig, err := GenerateSignature(token, key)
	if err != nil {
		return err
	}
	if sig == token.Signature {
		return nil
	}

	return errors.New("signature invalid")

}

/*GenerateSignature returns a signature based off of the data inside the token and the secret given
*
 */
func GenerateSignature(token *protos.SignedToken, key *os.File) (string, error) {
	var signData string
	signData += token.Header["encalg"] +
		token.Header["timealg"] +
		token.Header["type"] +
		token.Payload["username"] +
		token.Payload["expiration"]
	var sec []byte
	sec, err := ioutil.ReadFile(key.Name())
	if err != nil {
		return "", err
	}

	sig, err := generateHS256([]byte(signData), sec)
	return sig, err
}
