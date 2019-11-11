package servers

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	customtokens "MODE/servers/backend/networking/security/customtokens"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"errors"
	"io/ioutil"
	"net"
	"os"
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
func (serv *TLSserver) RequestToken(ctx context.Context, creds *protos.Credentials) (*protos.SignedToken, error) {
	head := &protos.SignedToken_Header{
		EncAlg:  "HS256",
		TimeAlg: "Unix",
		Type:    "MODE-TOKEN"}
	headData := head.EncAlg + head.TimeAlg + head.Type

	pay := &protos.SignedToken_Payload{
		Username:       creds.Username,
		ExpirationDate: time.Now().AddDate(0, 0, 7).Unix()}
	payData := pay.Username + string(pay.ExpirationDate)

	sec, err := ioutil.ReadFile(serv.privateKey.Name())
	if err != nil {
		return nil, errors.New("internal error: signature secret")
	}

	sig, err := customtokens.GenerateSignature([]byte(headData+payData), sec)
	if err != nil {
		return nil, errors.New("internal error: generating signature")
	}

	signature := &protos.SignedToken_Signature{
		Signature: sig}
	return &protos.SignedToken{
		Header: head, Payload: pay, Signature: signature}, nil

}
