package servers

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"MODE/servers/backend/networking/security/modesecurity"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//TLSserver replies to the essential authenticated RPC calls
type TLSserver struct {
	EssentialServer
	publicKey, privateKey string
	address               string
	port                  string
}

//NewTLSserver creates and returns a new TLSserver structure
func NewTLSserver(address, port string, publicKey, privateKey string) TLSserver {
	return TLSserver{address: address, port: port, publicKey: publicKey, privateKey: privateKey}
}

//Serve attaches to the specified address and port with the credentials given
func (serv *TLSserver) Serve() error {
	fmt.Printf("%v:%v", serv.address, serv.port)
	lis, err := net.Listen("tcp", serv.address+":"+serv.port)
	if err != nil {
		return err
	}
	creds, err := serv.getCreds(serv.publicKey, serv.privateKey)
	if err != nil {
		return err
	}
	serv.Server = grpc.NewServer(grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptors.TLSInterceptor),
		grpc.UnaryInterceptor(interceptors.EssentialInterceptor),
		grpc.ConnectionTimeout(5*time.Second),
		grpc.MaxConcurrentStreams(150000))
	generalservices.RegisterEssentialServer(serv.EssentialServer.Server, serv)
	generalservices.RegisterTokenSecurityServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
}

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

func (serv *TLSserver) getCreds(publicKey, privateKey string) (credentials.TransportCredentials, error) {
	if _, err := os.Stat(publicKey); os.IsNotExist(err) {
		return nil, errors.New("public key location invalid")
	}
	if _, err := os.Stat(privateKey); os.IsNotExist(err) {
		return nil, errors.New("private key location invalid")
	}
	creds, err := credentials.NewServerTLSFromFile(publicKey, privateKey)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func (serv *TLSserver) getPrivateKey() (*os.File, error) {
	if _, err := os.Stat(serv.privateKey); os.IsNotExist(err) {
		return nil, errors.New("private key location invalid")
	}
	return os.Open(serv.privateKey)
}
