package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "MODE/servers/backend/networking/proto/generated/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct{}

type picture struct {
	pictureContent []byte
	location       string
	textpost       string
	checksum       int64
}

//Serv is closed by outside function
var Serv grpc.Server

func main() {
	lis, err := net.Listen("tcp", ":3218")
	eCheck(err)

	Serv := grpc.NewServer()
	pb.RegisterMessagingServer(Serv, &server{})
	pb.RegisterHeartbeatingServer(Serv, &server{})
	pb.RegisterPictureUploadingServer(Serv, &server{})
	go waitForCloseCommand()
	if err := Serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *server) UploadPicture(stream pb.PictureUploading_UploadPictureServer) error {
	var returningStatusCode = pb.StatusCode_Failed
	var Pic picture
	var buffer bytes.Buffer
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			Pic.pictureContent = buffer.Bytes()
			if Pic.location == "" || Pic.textpost == "" || Pic.checksum == 0 {
				returningStatusCode = pb.StatusCode_Failed
			} else if picLen := len(Pic.pictureContent); picLen != int(Pic.checksum) {
				log.Printf("Piclen: %v   checksum: %v", picLen, Pic.checksum)
				returningStatusCode = pb.StatusCode_Failed
			} else {
				returningStatusCode = pb.StatusCode_Success
			}
			stream.SendAndClose(&pb.Status{Code: returningStatusCode})
			if returningStatusCode != pb.StatusCode_Success {
				return errors.New("exit code failed with unknown num")
			}
			Pic.writePictureToLocation()
			return nil

		}
		if err != nil {
			log.Println("Connection dropped, stopping upload")
			stream.SendAndClose(&pb.Status{Code: returningStatusCode})
			return errors.New("Connection Dropped")
		}
		switch chunkType := chunk.ChunkContent.(type) {
		case *pb.Chunk_Checksum:
			Pic.checksum = chunk.GetChecksum()
			log.Printf("Checksum is: %v", Pic.checksum)
		case *pb.Chunk_Location:
			Pic.location = chunk.GetLocation()
			log.Printf("Location for picture is: %v", Pic.location)
		case *pb.Chunk_TextPost:
			Pic.textpost = chunk.GetTextPost()
			log.Printf("Textpost is: %v", Pic.textpost)
		case *pb.Chunk_Content:
			length, err := buffer.Write(chunk.GetContent())
			if err != nil {
				panic(err)
			}
			log.Printf("Length of current picture buffer is: %v", length)
		default:
			return fmt.Errorf("Chunk.ChunkContent has unexpected type %T", chunkType)
		}
	}
}

func (s *server) SendMessage(ctx context.Context, in *pb.UserMessage) (*pb.Status, error) {
	log.Printf("Message: %v       Destination: %v", in.GetMessage(), in.GetDestination())
	return &pb.Status{Message: "Server received your message.", Code: pb.StatusCode_Success}, nil
}

func (s *server) UpdateInbox(mostRecentTime *pb.LastUpdateTime, stream pb.Messaging_UpdateInboxServer) error {
	return nil
}
func (s *server) ChangeMessageStatus(ctx context.Context, in *pb.UserMessage) (*pb.Status, error) {
	return &pb.Status{Message: "Server received your message.", Code: pb.StatusCode_Success}, nil
}

func (s *server) TradeHeartbeat(ctx context.Context, in *pb.Beat) (*pb.Status, error) {
	address, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Received beat from %v", address.Addr)
		return &pb.Status{Message: "Received beat.", Code: pb.StatusCode_Success}, nil
	}
	return nil, errors.New("Couldn't peer into context")

}

func closeServer() {
	Serv.GracefulStop()
}

func (Pic *picture) writePictureToLocation() error {
	f, err := os.Create(Pic.location)
	defer f.Close()
	if err != nil {
		return err
	}
	l, err := f.Write(Pic.pictureContent)

	if err != nil {
		return err
	}
	fmt.Println(l, "bytes written successfully")
	return nil
}

func waitForCloseCommand(server *grpc.Server) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "close" {
			closeServer()
		}
	}
	if scanner.Err() != nil {
		panic(err)
	}
}
