package interceptors

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"MODE/servers/backend/networking/security/modesecurity"
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
	timestr := "Request -- Time: " + time.Now().Format("2006-01-02 3:04:05PM") + "\tMethod: " + info.FullMethod
	now := time.Now()
	if info.FullMethod != "/protos.Essential/FetchCertificate" {
		c := make(chan error)
		go func(ch chan error) {
			fmt.Println("checking auth")
			ch <- authorize(ctx, info.FullMethod)
		}(c)
		auth := <-c
		authstr := "\tAuthorized: "
		if auth == nil {
			fmt.Println(timestr + authstr + "true")
		} else {
			fmt.Println(timestr + authstr + "false")
		}
	}
	// Calls the handler
	hC := make(chan interface{})
	errC := make(chan error)
	go func(hC chan interface{}, errC chan error) {
		h, err := handler(ctx, req)
		fmt.Printf("\nintercept took %v\n", time.Since(now))
		hC <- h
		errC <- err
	}(hC, errC)
	h := <-hC
	err := <-errC
	fmt.Println("returning from interrupt")
	return h, err
}

//Unimplemented authorize function for token-based auth
func authorize(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if md["password"] != nil {
		if md["password"][0] == "mypassword" {
			return nil
		}
	}
	if !ok {
		return errors.New("auth: metadata not found in context")
	}
	if md["type"] != nil && md["type"][0] == "mode-access-token" {
		if method != "/protos.TokenSecurity/RequestRefreshToken" &&
			method != "/protos.TokenSecurity/RequRequestAccessTokenestAccessToken" {
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
	tok, err := CreateTokenFromMD(md)
	if err != nil {
		return err
	}
	unixTimeStamp, err := strconv.ParseInt(tok.GetPayload()["expiration"], 10, 64)
	if err != nil {
		return errors.New("timestamp incorrect : " + tok.GetPayload()["expiration"])
	}
	if !time.Now().Before(time.Unix(unixTimeStamp, 0)) {
		return errors.New("auth: token has expired")
	}
	keyLoc := "/home/arline/go/src/MODE/servers/backend/certs/ModeKey.key"
	if _, err = os.Stat(keyLoc); err != nil {
		return err
	}
	err = modesecurity.ValidateToken(tok, keyLoc)
	if err != nil {
		return err
	}
	fmt.Printf("Token is valid for: %v\t", time.Unix(unixTimeStamp, 0).Sub(time.Now()))
	return nil
}

//AuthorizePassword compares the user/pass combo against the ones in the database
func authorizePassword(username, password string) error {

	//If username or password are missing, return false && unauthenticated error
	if username == "" || password == "" {
		return errors.New("auth: missing username &| password")
	}
	//Make call to database to compare
	//If the creds match, return true and nil error
	return nil
}

//CreateTokenFromMD returns a SignedToken created from the given metadata
func CreateTokenFromMD(md metadata.MD) (*generalservices.SignedToken, error) {
	if md["signature"] == nil {
		return &generalservices.SignedToken{}, errors.New("token is not signed")
	}
	return &generalservices.SignedToken{
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
	}, nil
}
