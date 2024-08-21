package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/dhuki/go-template/internal/infra/logger"
	"github.com/labstack/echo/v4"
)

type requestLog struct {
	Timestamp    time.Time   `json:"timestamp"`
	Method       string      `json:"method"`
	URL          string      `json:"url"`
	Status       int         `json:"status"`
	ResponseTime float64     `json:"responseTime"`
	ResponseSize int64       `json:"responseSize"`
	ReqBody      interface{} `json:"requestBody"`
}

func LogMiddleware() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			ctx := c.Request().Context()
			rl := requestLog{
				Timestamp:    start,
				Method:       c.Request().Method,
				URL:          c.Request().URL.RequestURI(),
				Status:       c.Response().Status,
				ResponseTime: float64(time.Since(start).Milliseconds()),
				ResponseSize: c.Response().Size,
				ReqBody:      ctx.Value(logger.MetaRequestBody),
			}
			logger.Info(ctx, "logger", "%+v", rl)
			return nil
		})
	})
}

func WithLogReqBody() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			req := c.Request()
			if req.Body != nil {
				reqBody, _ := io.ReadAll(req.Body)
				reqBody, _ = json.Marshal(strings.ReplaceAll(string(reqBody), " ", ""))
				ctx := req.Context()
				ctx = context.WithValue(ctx, logger.MetaRequestBody, string(reqBody))
				c.SetRequest(req.WithContext(ctx))
				req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		})
	})
}
