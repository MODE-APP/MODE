package servers

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"errors"
	"fmt"
	"net"
	"os"
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
		grpc.ConnectionTimeout(5*time.Second),
		grpc.MaxConcurrentStreams(150000))
	generalservices.RegisterTokenSecurityServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
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
