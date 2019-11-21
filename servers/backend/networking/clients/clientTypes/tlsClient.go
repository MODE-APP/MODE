package clients

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"context"
	"errors"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//TLSClient connects to a server using a TLS certificate
type TLSClient struct {
	EssentialClient
	generalservices.TokenSecurityClient
	cert         string
	AccessToken  *generalservices.SignedToken
	RefreshToken *generalservices.SignedToken
}

//NewTLSClient returns a TLS struct with the specified params
func NewTLSClient(address, port, certFile string) (TLSClient, error) {
	if _, err := os.Stat(certFile); err != nil {
		return TLSClient{}, err
	}
	tls := TLSClient{cert: certFile}
	tls.address = address
	tls.port = port
	return tls, nil
}

//Connect to the specified address defined in the EssentialClient struct
func (client *TLSClient) Connect() error {
	credentials, err := credentials.NewClientTLSFromFile(client.cert, "")
	if err != nil {
		panic(err)
	}
	client.ClientConn, err = grpc.Dial(client.address+":"+client.port, grpc.WithTransportCredentials(credentials))
	if err != nil {
		panic(err)
	}
	client.RegisterClientTypes()
	client.ctx, client.cancel = context.WithCancel(context.Background())
	return nil
}

//RegisterClientTypes registers the types of clients used by the TLS client
func (client *TLSClient) RegisterClientTypes() {
	client.EssentialClient.EssentialClient = generalservices.NewEssentialClient(client.ClientConn)
	client.TokenSecurityClient = generalservices.NewTokenSecurityClient(client.ClientConn)
}

//RequestRefreshToken asks the server for a refresh token and returns it
func (client *TLSClient) RequestRefreshToken(username, password string) (*generalservices.SignedToken, error) {
	creds := &generalservices.Credentials{
		Username: username, Password: password}
	client.ApplyCredentialsToMetadata(creds)
	return client.TokenSecurityClient.RequestRefreshToken(client.ctx, creds)

}

//RequestAccessToken asks the server for an access token and returns it
func (client *TLSClient) RequestAccessToken() (*generalservices.SignedToken, error) {
	return client.TokenSecurityClient.RequestAccessToken(client.ctx, client.RefreshToken)

}

//ApplyTokenToMetadata applies the client's token to the client's metadata
func (client *TLSClient) ApplyTokenToMetadata(token *generalservices.SignedToken) error {
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

func (client *TLSClient) ApplyCredentialsToMetadata(creds *generalservices.Credentials) error {
	md, _ := metadata.FromOutgoingContext(client.ctx)
	if md["password"] == nil {
		client.ctx = metadata.AppendToOutgoingContext(client.ctx, "password", creds.GetPassword(), "username", creds.GetUsername())
	}
	return nil
}
