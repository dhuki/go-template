package health

import (
	"context"

	proto "github.com/dhuki/go-template/internal/adapter/grpc/v1/pb"
)

type Service interface {
	HealthCheck(ctx context.Context) (err error)
	HealthCheckGRPC(ctx context.Context) (req *proto.HealthCheckAPIResponse, err error)
}

type Repository interface {
	Ping(ctx context.Context) (err error)
}
