package clients

import (
	"MODE/servers/backend/networking/proto/generated/protos"
	proto "MODE/servers/backend/networking/proto/generated/protos"
	"MODE/servers/backend/networking/security/customtokens"
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
	cert  *os.File
	Token customtokens.BasicToken
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

func (client *TLSClient) RequestToken(username, password string) (customtokens.BasicToken, error) {
	token, err := client.TokenSecurityClient.RequestToken(client.ctx, &protos.Credentials{
		Username: username, Password: password})
	return customtokens.BasicToken{
		TokenHeader:    token.GetHeader().GetAlg() + token.GetHeader().GetType(),
		TokenPayload:   token.GetPayload().GetUsername(),
		TokenSignature: token.GetSignature().GetSignature()}, err
}

func (client *TLSClient) ApplyTokenToMetadata() error {
	md := metadata.New(map[string]string{"tokenheader": client.Token.TokenHeader,
		"tokenpayload":   client.Token.TokenPayload,
		"tokensignature": client.Token.TokenSignature})
	if (client.Token == customtokens.BasicToken{}) {
		return errors.New("could not create md from token")
	}
	client.ctx = metadata.NewOutgoingContext(client.ctx, md)
	return nil
}
