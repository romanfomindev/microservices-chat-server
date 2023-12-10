package interceptor

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func StreamAccessInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("start StreamAccessInterceptor")

	accessToken, err := GetAccessToken(stream.Context())
	if err != nil {
		return err
	}

	email, err := CheckAccess(accessToken, info.FullMethod)
	log.Printf("response: %+v %+v", email, err)
	if err != nil {
		return err
	}

	return handler(srv, &grpc_middleware.WrappedServerStream{
		ServerStream:   stream,
		WrappedContext: context.WithValue(context.Background(), "email", email),
	})
}
