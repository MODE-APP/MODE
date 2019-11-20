package servers

import (
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	interceptors "MODE/servers/backend/networking/servers/interceptorTypes"
	"context"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

//EssentialServer replies to the essential unauthenticated RPC calls
type EssentialServer struct {
	*grpc.Server
	generalservices.EssentialServer
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
	generalservices.RegisterEssentialServer(serv.Server, serv)
	err = serv.Server.Serve(lis)
	return err
}

func (*EssentialServer) TestCall(ctx context.Context, info *generalservices.Info) (*generalservices.GeneralStatus, error) {
	return &generalservices.GeneralStatus{
		Status: generalservices.GeneralStatus_OK}, nil
}

//FetchCertificate returns a copy of the public key to the client requesting it
func (*EssentialServer) FetchCertificate(ctx context.Context, info *generalservices.Info) (*generalservices.File, error) {
	buf, err := ioutil.ReadFile("/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt")
	if err != nil {
		log.Println("err wasnt nil")
		return nil, err
	}
	log.Println("found cert successfully")
	return &generalservices.File{
		FileBytes: buf, FileName: "MODE_CERT.crt", Checksum: 0}, nil
}
