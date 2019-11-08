package servers

import (
	proto "MODE/servers/backend/networking/proto/generated/protos"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
)

//EssentialServer replies to the essential unauthenticated RPC calls
type EssentialServer struct {
	*grpc.Server
	proto.EssentialServer
	address string
	port    string
}

//NewEssentialServer creates and returns a new EssentialServer object
func NewEssentialServer(address, port string) EssentialServer {
	return EssentialServer{address: address, port: port}
}

//Serve attaches to the specified address and port as a basic EssentialServer
func (serv *EssentialServer) Serve() error {
	lis, err := net.Listen("tcp", serv.address+":"+serv.port)
	if err != nil {
		return err
	}
	serv.Server = grpc.NewServer(grpc.UnaryInterceptor(interceptors.EssentialInterceptor))
	proto.RegisterEssentialServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
}

//FetchCertificate returns a copy of the public key to the client requesting it
func (*EssentialServer) FetchCertificate(ctx context.Context, info *proto.Info) (*proto.File, error) {
	buf, err := ioutil.ReadFile("/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt")
	if err != nil {
		return nil, err
	}
	return &proto.File{
		FileBytes: buf, FileName: "MODE_CERT.crt", Checksum: 0}, nil
}
