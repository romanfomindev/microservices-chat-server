package interceptor

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/romanfomindev/microservices-chat-server/internal/client/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const AUTH_PREFIX = "Bearer "

func AccessInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("start AccessInterceptor")
	accessToken, err := GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	email, err := CheckAccess(accessToken, info.FullMethod)
	log.Printf("response: %+v %+v", email, err)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "email", email)
	res, err := handler(ctx, req)

	return res, err
}

func GetAccessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], AUTH_PREFIX) {
		return "", errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], AUTH_PREFIX)

	return accessToken, nil
}

func CheckAccess(accessToken, method string) (string, error) {
	ctx := context.Background()

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", auth.ServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to dial GRPC client: %v", err)
		return "", err
	}

	client := auth.NewAuth(conn)

	return client.CheckAccess(ctx, method)
}
