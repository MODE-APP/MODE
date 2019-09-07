package main

import (
	pb "MODE/servers/backend/networking/proto/generated"
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct{}
type picture struct {
	file     *os.File
	location string
	textpost string
	checksum int64
}

const BUFFER_SIZE int = 4096

type Client struct {
	pb.HeartbeatingClient
	pb.MessagingClient
	pb.PictureUploadingClient
}

func main() {
	conn, err := grpc.Dial("73.83.1.188:3218", grpc.WithInsecure())
	eCheck(err)
	defer conn.Close()
	client := NewClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	file, err := os.Open("/home/arline/Downloads/testy.jpg")
	if err != nil {
		panic(err)
	}
	pic := picture{file, "C:/Users/chase/Pictures/testPic.jpg", "posty test", simpleChecksum(file)}
	startTime := time.Now()
	client.sendPicture(ctx, pic)
	endTime := time.Since(startTime)
	log.Printf("Time elapsed: %v", endTime)
}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func NewClient(conn *grpc.ClientConn) Client {
	return Client{pb.NewHeartbeatingClient(conn), pb.NewMessagingClient(conn), pb.NewPictureUploadingClient(conn)}
}

func (client *Client) sendPicture(ctx context.Context, Picture picture) {
	buf := make([]byte, BUFFER_SIZE)
	uploader, err := client.UploadPicture(ctx)
	if err != nil {
		panic(err)
	}
	for n, err := Picture.file.Read(buf); err != io.EOF; n, err = Picture.file.Read(buf) {
		uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Content{buf[0:n]}})
	}
	log.Printf("out of file to read")
	uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Location{Picture.location}})
	uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_TextPost{Picture.textpost}})
	uploader.Send(&pb.Chunk{ChunkContent: &pb.Chunk_Checksum{Picture.checksum}})

	status, err := uploader.CloseAndRecv()
	if err != nil && err != io.EOF {
		panic(err)
	}
	log.Printf("Picture upload status message: %v     Code: %v", status.GetMessage(), status.GetCode())
}

func simpleChecksum(file *os.File) int64 {
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return fileStat.Size()
}
