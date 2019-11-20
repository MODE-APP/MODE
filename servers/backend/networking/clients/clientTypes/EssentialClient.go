package clients

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

//EssentialClient is the bare essentials to interact with the server
type EssentialClient struct {
	address string
	port    string
	generalservices.EssentialClient
	*grpc.ClientConn
	cancel context.CancelFunc
	ctx    context.Context
}

//IsConnected returns if the connection is ready to send data
func (client *EssentialClient) IsConnected() (connected bool) {
	if client.ClientConn != nil {
		return client.GetState() == connectivity.Ready
	}
	return false
}

//NewEssentialClient generates a new EssentialClient struct with a nil connection
func NewEssentialClient(address, port string) EssentialClient {
	return EssentialClient{address: address, port: port}
}

//Connect connects the client using unsecure transport - do not use for secure transfer
func (client *EssentialClient) Connect() error {
	var err error
	client.ClientConn, err = grpc.Dial(client.address+":"+client.port, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client.EssentialClient = generalservices.NewEssentialClient(client.ClientConn)
	client.ctx, client.cancel = context.WithCancel(context.Background())
	return err
}

//FetchCertificate grabs public key from the server and is returned
func (client *EssentialClient) FetchCertificate() (fileBuf []byte, filename string, err error) {
	if client.ClientConn != nil && client.ctx != nil && client.EssentialClient != nil {
		file, err := client.EssentialClient.FetchCertificate(client.ctx, &generalservices.Info{})
		if err != nil {
			return nil, "", err
		}
		log.Println("returning cert")
		return file.FileBytes, file.FileName, nil
	}
	return nil, "", errors.New("connection not ready")
}

func (client *EssentialClient) RegisterClientTypes() {
	client.EssentialClient = generalservices.NewEssentialClient(client.ClientConn)
}

func (client *EssentialClient) TestCall() {
	client.EssentialClient.TestCall(client.ctx, &generalservices.Info{})
}
