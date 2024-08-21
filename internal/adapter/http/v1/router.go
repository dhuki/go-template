package v1

import (
	"context"
	"fmt"

	httpHandler "github.com/dhuki/go-template/internal/adapter/http"
	midw "github.com/dhuki/go-template/internal/adapter/http/middleware"
	"github.com/dhuki/go-template/internal/infra/configloader"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type svc struct {
	e       *echo.Echo
	addr    string
	handler httpHandler.Handler
}

func NewHttpHandlerV1(handler httpHandler.Handler, port int) *svc {
	e := echo.New()
	e.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Timeout:      configloader.Conf.App.Timeout,
		ErrorMessage: "timeout",
	}))
	e.Use(midw.CollectMetadata())
	e.Use(echoMiddleware.Recover())
	e.Use(midw.LogMiddleware())

	return &svc{e: e, addr: fmt.Sprintf(":%d", port), handler: handler}
}

func (s *svc) Start(ctx context.Context) error {
	v1Group := s.e.Group("/api/v1")
	s.RegistHealthRoute(v1Group)
	return s.e.Start(s.addr)
}

func (s *svc) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
