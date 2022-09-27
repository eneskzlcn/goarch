package router

import (
	"context"
	"net/http"
)

type MethodHandlers map[string]http.HandlerFunc

type Router interface {
	Handle(pattern string, methodHandlers MethodHandlers)
	ExtractURLParam(ctx context.Context, key string) string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
