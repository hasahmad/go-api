package context

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type httpRequestIdKey string

const requestIdKey = httpRequestIdKey("X-REQUEST-ID")

func SetRequestID(ctx context.Context, reqID string) context.Context {
	if reqID == "" {
		return ctx
	}

	if ctx.Value(requestIdKey) != nil {
		return ctx
	}

	return context.WithValue(ctx, requestIdKey, reqID)
}

func GetRequestID(ctx context.Context) string {
	if requestID := ctx.Value(requestIdKey); requestID != nil {
		return fmt.Sprintf("%s", requestID)
	}

	return uuid.NewString()
}
