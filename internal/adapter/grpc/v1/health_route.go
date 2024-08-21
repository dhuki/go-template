package v1

import (
	"context"

	proto "github.com/dhuki/go-template/internal/adapter/grpc/v1/pb"
)

func (svc *svc) HealthCheckAPI(ctx context.Context, req *proto.HealthCheckAPIRequest) (*proto.HealthCheckAPIResponse, error) {
	resp, err := svc.handler.HealthService.HealthCheckGRPC(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
