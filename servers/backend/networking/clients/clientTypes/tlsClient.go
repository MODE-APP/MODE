package clients

import (
	proto "MODE/servers/backend/networking/proto/generated/protos"
	"context"
	"errors"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//TLSClient connects to a server using a TLS certificate
type TLSClient struct {
	EssentialClient
	proto.TokenSecurityClient
	cert *os.File
}

//NewTLSClient returns a TLS struct with the specified params
func NewTLSClient(address, port string, certFile *os.File) (TLSClient, error) {
	if _, err := os.Stat(certFile.Name()); err != nil {
		return TLSClient{}, err
	}
	tls := TLSClient{cert: certFile}
	tls.address = address
	tls.port = port
	return tls, nil
}

//Connect to the specified address defined in the EssentialClient struct
func (client *TLSClient) Connect() error {
	if client.cert == nil {
		return errors.New("Cert was nil")
	}
	log.Printf("Cert name: %v", client.cert.Name())
	credentials, err := credentials.NewClientTLSFromFile(client.cert.Name(), "")
	if err != nil {
		return err
	}
	client.ClientConn, err = grpc.Dial(client.address+":"+client.port, grpc.WithTransportCredentials(credentials))
	if err != nil {
		return err
	}
	client.RegisterClientTypes()
	client.ctx, client.cancel = context.WithCancel(context.Background())
	return nil
}

//RegisterClientTypes registers the types of clients used by the TLS client
func (client *TLSClient) RegisterClientTypes() {
	client.EssentialClient.RegisterClientTypes()
	client.TokenSecurityClient = proto.NewTokenSecurityClient(client.ClientConn)
}
