package v1

import (
	"net/http"

	httpHelper "github.com/dhuki/go-template/internal/adapter/http"
	midw "github.com/dhuki/go-template/internal/adapter/http/middleware"
	"github.com/dhuki/go-template/internal/infra/logger"
	"github.com/labstack/echo/v4"
)

func (s *svc) RegistHealthRoute(app *echo.Group) {
	v1GroupHealth := app.Group("/health")

	v1GroupHealth.GET("", s.healthCheck())
	v1GroupHealth.POST("/example", s.example(), midw.WithLogReqBody())
}

func (s *svc) healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if err := s.handler.HealthService.HealthCheck(ctx); err != nil {
			return httpHelper.ResponseError(c, http.StatusInternalServerError, err)
		}
		return httpHelper.ResponseSuccess(c, "success", nil)
	}
}

func (s *svc) example() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		logger.Info(ctx, "success call function example", "example handler")
		return httpHelper.ResponseSuccess(c, "success", nil)
	}
}
