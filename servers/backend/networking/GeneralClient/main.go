package main

import (
	pb "MODETEST/proto"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

const BUFFER_SIZE int = 2048

type client struct {
	pb.HeartbeatingClient
	pb.MessagingClient
	pb.PictureUploadingClient
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	eCheck(err)
	defer conn.Close()

}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func NewHeartbeatMessenger(conn *grpc.ClientConn) client {
	return client{pb.NewHeartbeatingClient(conn), pb.NewMessagingClient(conn), pb.NewPictureUploadingClient(conn)}
}

func sendPicture(uploader pb.PictureUploading_UploadPictureClient, file *os.File, location string, textpost string) {
	buf := make([]byte, 4096)
	for {
		if n, err := file.Read(buf); err != io.EOF {
			if n > 0 {
				uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Content{buf[0:n]}})
				log.Printf("last 10 bytes")
			} else {
				log.Printf("READING 0")
			}
		} else {
			log.Printf("out of file to read")
			fileStat, err := file.Stat()
			if err != nil {
				panic(err)
			}
			uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Location{location}})
			uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_TextPost{textpost}})
			uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Checksum{fileStat.Size()}})

			status, err := uploader.CloseAndRecv()
			if err != nil {
				panic(err)
			}
			log.Printf("Picture upload status message: %v     Code: %v", status.GetMessage(), status.GetCode())
		}

	}

}
