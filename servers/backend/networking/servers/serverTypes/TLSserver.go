package servers

import (
	proto "MODE/servers/backend/networking/proto/generated/protos"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"errors"
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
	proto.RegisterEssentialServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
}
