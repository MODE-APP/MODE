package interceptors

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	"MODE/servers/backend/networking/security/customtokens"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//TLSInterceptor checks the given username/password username/token combination against the ones in the database
func TLSInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authorize when fetching certificate/refreshtoken
	fmt.Printf("Request -- Time: %v\t  Method: %s\t", time.Now().Format("2006-01-02 3:04:05PM"), info.FullMethod)
	if info.FullMethod != "/protos.Essential/FetchCertificate" &&
		info.FullMethod != "/protos.TokenSecurity/RequestRefreshToken" {
		if err := authorize(ctx, info.FullMethod); err != nil {
			fmt.Printf("Authorized: %v", false)
			fmt.Println()
			return nil, err
		}
		fmt.Printf("Authorized: %v", true)

	}
	fmt.Println()
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

//Unimplemented authorize function for token-based auth
func authorize(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if md["username"] != nil {
		fmt.Printf("Username: %v\t", md["username"][0])
	}
	if !ok {
		return errors.New("auth: metadata not found in context")
	}
	if md["type"] != nil && md["type"][0] == "mode-access-token" {
		if method != "/protos.TokenSecurity/RequestRefreshToken" &&
			method != "/protos.TokenSecurity/RequestAccessToken" {
			return authorizeToken(md)
		}
		fmt.Printf("Token is of wrong type\t")
		return errors.New("auth: wrong token type for method call")
	} else if md["type"] != nil && md["type"][0] == "mode-refresh-token" {
		if method == "/protos.TokenSecurity/RequestAccessToken" {
			return authorizeToken(md)
		}
		fmt.Printf("Token is of wrong type\t")
		return errors.New("auth: wrong token type for method call")
	}
	return errors.New("auth: credentials missing")
}

//AuthorizeToken generates the hash and compares the result against the token data given
func authorizeToken(md metadata.MD) error {
	tok := &protos.SignedToken{
		Header: map[string]string{
			"encalg":  md["encalg"][0],
			"timealg": md["timealg"][0],
			"type":    md["type"][0],
		},
		Payload: map[string]string{
			"username":   md["username"][0],
			"expiration": md["expiration"][0],
		},
		Signature: md["signature"][0],
	}

	unixTimeStamp, err := strconv.ParseInt(tok.GetPayload()["expiration"], 10, 64)
	if err != nil {
		return errors.New("timestamp incorrect : " + tok.GetPayload()["expiration"])
	}
	if !time.Now().Before(time.Unix(unixTimeStamp, 0)) {
		return errors.New("auth: token has expired")
	}
	key, err := os.Open("/home/arline/go/src/MODE/servers/backend/certs/ModeKey.key")
	if err != nil {
		return err
	}
	if _, err = os.Stat(key.Name()); err != nil {
		return err
	}
	err = customtokens.ValidateToken(tok, key)
	if err == nil {
		fmt.Printf("Token is valid for: %v\t", time.Unix(unixTimeStamp, 0).Sub(time.Now()))
	}
	return err
}

//AuthorizePassword compares the user/pass combo against the ones in the database
func authorizePassword(username, password string) error {
	//If username or password are missing, return false && unauthenticated error
	if username == "" || password == "" {
		errors.New("auth: missing username &| password")
	}
	//Make call to database to compare
	//If the creds match, return true and nil error
	return nil
}
