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
	Token customtokens.Token
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

//RequestToken asks the server a token and returns it
func (client *TLSClient) RequestToken(username, password string) (customtokens.Token, error) {
	token, err := client.TokenSecurityClient.RequestToken(client.ctx, &protos.Credentials{
		Username: username, Password: password})
	if err == nil {
		return customtokens.Token{
			TokenHeader: map[string]string{
				"encalg":  token.GetHeader().GetEncAlg(),
				"timealg": token.GetHeader().GetTimeAlg(),
				"type":    token.GetHeader().GetType(),
			},
			TokenPayload: map[string]string{
				"username":   token.GetPayload().GetUsername(),
				"expiration": string(token.GetPayload().GetExpirationDate()),
			},
			TokenSignature: token.GetSignature().GetSignature(),
		}, nil
	}
	return customtokens.Token{}, err
}

//ApplyTokenToMetadata applies the client's token to the client's metadata
func (client *TLSClient) ApplyTokenToMetadata() error {
	tok := client.Token
	if tok.TokenSignature == "" {
		return errors.New("auth: missing signature on token")
	}
	tokMap := map[string]string{}
	for k, v := range tok.TokenHeader {
		if tokMap[k] == "" {
			tokMap[k] = v
		}
	}
	for k, v := range tok.TokenPayload {
		if tokMap[k] == "" {
			tokMap[k] = v
		}
	}
	tokMap["signature"] = tok.TokenSignature
	md := metadata.New(tokMap)
	client.ctx = metadata.NewOutgoingContext(client.ctx, md)
	return nil
}
