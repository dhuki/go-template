package grpc

import (
	"github.com/dhuki/go-template/internal/core/health"
	"github.com/dhuki/go-template/internal/infra/database/repository"
)

type Handler struct {
	HealthService health.Service
}

func NewHandler(repo repository.IRepository) Handler {
	return Handler{
		HealthService: health.NewService(repo, nil),
	}
}
