package clients

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	proto "MODE/servers/backend/networking/proto/generated/protos"
	"context"
	"errors"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//TLSClient connects to a server using a TLS certificate
type TLSClient struct {
	EssentialClient
	proto.TokenSecurityClient
	cert         *os.File
	AccessToken  *protos.SignedToken
	RefreshToken *protos.SignedToken
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

//RequestRefreshToken asks the server for a refresh token and returns it
func (client *TLSClient) RequestRefreshToken(username, password string) (*protos.SignedToken, error) {
	return client.TokenSecurityClient.RequestRefreshToken(client.ctx, &protos.Credentials{
		Username: username, Password: password})

}

//RequestAccessToken asks the server for an access token and returns it
func (client *TLSClient) RequestAccessToken() (*protos.SignedToken, error) {
	return client.TokenSecurityClient.RequestAccessToken(client.ctx, client.RefreshToken)

}

//ApplyTokenToMetadata applies the client's token to the client's metadata
func (client *TLSClient) ApplyTokenToMetadata(token *protos.SignedToken) error {
	if token.Signature == "" {
		return errors.New("auth: missing signature on token")
	}
	tokMap := map[string]string{}
	for k, v := range token.Header {
		if tokMap[k] == "" {
			tokMap[k] = v
		}
	}
	for k, v := range token.Payload {
		if tokMap[k] == "" {
			tokMap[k] = v
		}
	}
	tokMap["signature"] = token.Signature
	md := metadata.New(tokMap)
	client.ctx = metadata.NewOutgoingContext(client.ctx, md)
	return nil
}
