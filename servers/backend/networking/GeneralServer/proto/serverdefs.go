package serverdefs

import (
	pb "MODE/servers/backend/networking/proto/generated/protos"
)

//HeartBeatingServer manages the tracking of disconnected users
type HeartBeatingServer struct {
	pb.HeartbeatingServer
}

//MessagingServer manages the messaging between users/groupchats
type MessagingServer struct {
	pb.MessagingServer
}

//CommentingServer manages the comments on posts
type CommentingServer struct {
	pb.CommentingServer
}

//PictureUploadingServer manages the uploading of pictures to various locations (including groupchats/messages)
type PictureUploadingServer struct {
	pb.PictureUploadingServer
}

//BetaServer will handle all of the functions on one machine instead of being distributed
type BetaServer struct {
	PictureUploadingServer
	CommentingServer
	MessagingServer
	HeartBeatingServer
}

//NewBetaServer returns a BetaServer that has not started yet
func NewBetaServer() *BetaServer {

}
