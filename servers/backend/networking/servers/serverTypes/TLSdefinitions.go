package servers

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"MODE/servers/backend/networking/security/modesecurity"
	"context"
	"fmt"
	"strconv"
	"time"
)

//RequestAccessToken creates an access token that can be used as authentication
func (serv *TLSserver) RequestAccessToken(ctx context.Context, refreshToken *generalservices.SignedToken) (*generalservices.SignedToken, error) {
	refreshToken.Header["type"] = "mode-access-token"
	refreshToken.Payload["expiration"] = strconv.FormatInt(time.Now().Add(time.Minute*30).Unix(), 10)
	sig, err := modesecurity.GenerateSignature(refreshToken, serv.privateKey)
	if err != nil {
		return &generalservices.SignedToken{}, err
	}
	refreshToken.Signature = sig
	return refreshToken, nil
}

//RequestRefreshToken creates a refresh token that can be used to obtain access tokens
func (serv *TLSserver) RequestRefreshToken(ctx context.Context, creds *generalservices.Credentials) (*generalservices.SignedToken, error) {
	//check database for credentials, return error if not found
	//continue if found
	tok := &generalservices.SignedToken{
		Header: map[string]string{
			"encalg":  "hs256",
			"timealg": "unix",
			"type":    "mode-refresh-token",
		},
		Payload: map[string]string{
			"username":   creds.GetUsername(),
			"expiration": strconv.FormatInt(time.Now().AddDate(0, 0, 1).Unix(), 10),
		},
	}
	sig, err := modesecurity.GenerateSignature(tok, serv.privateKey)
	if err != nil {
		fmt.Printf("err: %v", err)
		return &generalservices.SignedToken{}, err
	}
	tok.Signature = sig
	return tok, nil
}
