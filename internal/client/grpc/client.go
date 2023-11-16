package grpc

import "context"

type Auth interface {
	CheckAccess(ctx context.Context, endpoint string) error
}
