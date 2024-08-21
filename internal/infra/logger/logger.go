package logger

import (
	"context"

	log "github.com/sirupsen/logrus"
)

const (
	MetaCorrelationID string = "X-Correlation-Id"
	MetaRequestBody   string = "X-Req-Body"
)

func Info(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Infof(format, args...)
}

func Warn(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Warnf(format, args...)
}

func Debug(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Debugf(format, args...)
}

func Error(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Errorf(format, args...)
}

func Fatal(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Fatalf(format, args...)
}

func getEntry(ctx context.Context, ctxName string) *log.Entry {
	return log.WithFields(log.Fields{
		"context":       ctxName,
		"correlationId": ctx.Value(MetaCorrelationID),
	})
}
