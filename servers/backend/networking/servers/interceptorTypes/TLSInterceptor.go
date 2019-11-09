package interceptors

import (
	"MODE/servers/backend/networking/security/customtokens"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//TLSInterceptor checks the given username/password username/token combination against the ones in the database
func TLSInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Skip authorize when fetching certificate/tokens
	if info.FullMethod != "/protos.Essential/FetchCertificate" &&
		info.FullMethod != "/protos.TokenSecurity/RequestToken" {
		if ok, err := authorize(ctx); !ok {
			log.Printf("Request - Method: %s\tTime: %s\tAuthorized: %s", info.FullMethod, time.Now(), "false")
			return nil, err
		}
	}
	log.Printf("Request - Method: %s\tTime: %s\tAuthorized: %s", info.FullMethod, time.Now(), "true")
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

//Unimplemented authorize function for token-based auth
func authorize(ctx context.Context) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("auth: metadata not found in rpc")
	}
	if md["password"] != nil {
		ok, err := AuthorizePassword(md["username"][0], md["password"][0])
		if !ok || err != nil {
			return false, errors.New("auth: password could not be authenticated")
		}
	} else if md["tokenheader"] != nil {
		tokenHeader := md["tokenheader"]
		tokenPayload := md["tokenpayload"]
		tokenSignature := md["tokensignature"]
		ok, err := AuthorizeToken(strings.Join(tokenHeader, ""), strings.Join(tokenPayload, ""), tokenSignature[0])
		return ok, err
	}
	return false, errors.New("auth: no credentials given")
}

//AuthorizeToken generates the hash and compares the result against the token data given
func AuthorizeToken(tokenHeader, tokenPayload, tokenSignature string) (bool, error) {
	sec, err := ioutil.ReadFile("/home/arline/go/src/MODE/servers/backend/certs/ModeKey.key")
	if err != nil {
		return false, errors.New("internal error, key location incorrect")
	}

	localSig, err := customtokens.GenerateSignature([]byte(tokenHeader+tokenPayload), sec)
	if err != nil {
		return false, errors.New("auth: could not sign token data")
	}
	log.Printf("\nlocalsig: %v\ntokenSig: %v", localSig, tokenSignature)
	if localSig == tokenSignature {
		return true, nil
	}
	return false, errors.New("auth: token not recognized")
}

//AuthorizePassword compares the user/pass combo against the ones in the database
func AuthorizePassword(username, password string) (bool, error) {
	//If username or password are missing, return false && unauthenticated error
	if username == "" || password == "" {
		return false, errors.New("auth: missing username &| password")
	}
	//Make call to database to compare
	//If the creds match, return true and nil error
	return true, nil
}
