package servers

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	"MODE/servers/backend/networking/security/customtokens"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"errors"
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
	publicKey  *os.File
	privateKey *os.File
	address    string
	port       string
}

//NewTLSserver creates and returns a new TLSserver structure
func NewTLSserver(address, port string, publicKey, privateKey *os.File) TLSserver {
	return TLSserver{address: address, port: port, publicKey: publicKey, privateKey: privateKey}
}

//Serve attaches to the specified address and port with the credentials given
func (serv *TLSserver) Serve() error {
	if _, err := os.Stat(serv.publicKey.Name()); os.IsNotExist(err) {
		return errors.New("public key location invalid")
	}
	if _, err := os.Stat(serv.privateKey.Name()); os.IsNotExist(err) {
		return errors.New("private key location invalid")
	}
	creds, err := credentials.NewServerTLSFromFile(serv.publicKey.Name(), serv.privateKey.Name())
	if err != nil {
		return errors.New("invalid credential creation")
	}
	lis, err := net.Listen("tcp", serv.address+":"+serv.port)
	if err != nil {
		return err
	}
	serv.Server = grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptors.TLSInterceptor))
	protos.RegisterEssentialServer(serv.Server, serv)
	protos.RegisterTokenSecurityServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
}

//RequestToken creates a token that can be used to verify identity
func (serv *TLSserver) RequestAccessToken(ctx context.Context, refreshToken *protos.SignedToken) (*protos.SignedToken, error) {
	refreshToken.Header["type"] = "mode-access-token"
	refreshToken.Payload["expiration"] = strconv.FormatInt(time.Now().Add(time.Minute*30).Unix(), 10)
	sig, err := customtokens.GenerateSignature(refreshToken, serv.privateKey)
	if err != nil {
		return &protos.SignedToken{}, err
	}
	refreshToken.Signature = sig
	return refreshToken, nil
}

func (serv *TLSserver) RequestRefreshToken(ctx context.Context, creds *protos.Credentials) (*protos.SignedToken, error) {
	//check database for credentials, return error if not found
	//continue if found
	tok := &protos.SignedToken{
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
	sig, err := customtokens.GenerateSignature(tok, serv.privateKey)
	if err != nil {
		return &protos.SignedToken{}, err
	}
	tok.Signature = sig
	return tok, nil
}
