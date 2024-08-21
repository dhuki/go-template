package health

import (
	"context"
	"fmt"

	proto "github.com/dhuki/go-template/internal/adapter/grpc/v1/pb"
	"github.com/dhuki/go-template/internal/infra/logger"
)

type service struct {
	repoPg    Repository
	repoMysql Repository
}

func NewService(healthRepoPg Repository, healthRepoMysql Repository) Service {
	return &service{
		repoPg:    healthRepoPg,
		repoMysql: healthRepoMysql,
	}
}

func (s *service) HealthCheck(ctx context.Context) error {
	ctxName := fmt.Sprintf("%T.HealthCheck", s)
	if err := s.repoPg.Ping(ctx); err != nil {
		logger.Error(ctx, ctxName, "u.repository.Ping, got err: %v", err)
		return err
	}
	return nil
}

func (s *service) HealthCheckGRPC(ctx context.Context) (*proto.HealthCheckAPIResponse, error) {
	ctxName := fmt.Sprintf("%T.HealthCheckGRPC", s)
	if err := s.repoPg.Ping(ctx); err != nil {
		logger.Error(ctx, ctxName, "u.repository.Ping, got err: %v", err)
		return nil, err
	}
	return &proto.HealthCheckAPIResponse{
		Message: "Success",
	}, nil
}
