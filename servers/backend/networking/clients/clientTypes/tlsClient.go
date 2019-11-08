package clients

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//TLSClient connects to a server using a TLS certificate
type TLSClient struct {
	EssentialClient
	certification os.File
}

//Connect to the specified address defined in the EssentialClient struct
func (client *TLSClient) Connect() error {
	credentials, err := credentials.NewClientTLSFromFile(client.certification.Name(), "")
	if err != nil {
		return err
	}
	client.ClientConn, err = grpc.Dial(client.address+":"+client.port, grpc.WithTransportCredentials(credentials))
	if err != nil {
		return err
	}
	return nil
}
