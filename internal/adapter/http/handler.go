package http

import (
	"github.com/dhuki/go-template/internal/core/health"
	"github.com/dhuki/go-template/internal/infra/database/repository"
)

type Handler struct {
	HealthService health.Service
}

func NewHandler(repoPG, repoMysql repository.IRepository) Handler {
	return Handler{
		HealthService: health.NewService(repoPG, repoMysql),
	}
}
