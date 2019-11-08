package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

//TLSInterceptor checks the given username/password username/token combination against the ones in the database
func TLSInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Skip authorize when fetching certificate/tokens
	if info.FullMethod != "/proto.EssentialService/FetchCertificate" &&
		info.FullMethod != "/proto.TokenSecurityService/FetchToken" {
		if err := authorize(ctx); err != nil {
			log.Printf("Request - Method: %s\tTime: %s\tAuthorized: %s", info.FullMethod, time.Now(), "false")
			return nil, err
		}
	}
	log.Printf("Request - Method: %s\tTime: %s\tAuthorized: %s", info.FullMethod, time.Now(), "true")
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

//Unimplemented authorize function for token-based auth
func authorize(ctx context.Context) error {
	return nil
}
