package middleware

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CollectMetadata() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			// get correlation id
			corrId := req.Header.Get(string(echo.HeaderXCorrelationID))
			if corrId == "" {
				corrId = uuid.NewString()
			}

			ctx = context.WithValue(ctx, echo.HeaderXCorrelationID, corrId)
			c.SetRequest(req.WithContext(ctx))
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
