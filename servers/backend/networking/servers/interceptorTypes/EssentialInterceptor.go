package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

//EssentialInterceptor is used to check users against the Essential service RPC
func EssentialInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	//Prints out request at the specified time
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}
