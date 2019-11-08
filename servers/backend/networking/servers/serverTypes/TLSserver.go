package servers

import (
	protos "MODE/servers/backend/networking/proto/generated/protos"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"os"

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
func (serv *TLSserver) RequestToken(ctx context.Context, creds *protos.Credentials) (*protos.SignedToken, error) {
	head := &protos.SignedToken_Header{
		Alg:  "HS256",
		Type: "MODE-TOKEN"}
	headData := head.Alg + head.Type

	pay := &protos.SignedToken_Payload{
		Username: creds.Username}

	sec, err := ioutil.ReadFile(serv.privateKey.Name())
	if err != nil {
		return nil, err
	}

	sig, err := GenerateSignature([]byte(headData+pay.Username), sec)
	signature := &protos.SignedToken_Signature{
		Signature: sig}
	return &protos.SignedToken{
		Header: head, Payload: pay, Signature: signature}, nil

}

//GenerateSignature returns a HS256 signature used for tokenization
func GenerateSignature(data, secret []byte) (string, error) {
	h := hmac.New(sha256.New, secret)
	_, err := h.Write(data)
	if err != io.ErrClosedPipe {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
