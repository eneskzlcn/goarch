package router

import (
	"context"
	"github.com/nicolasparada/go-mux"
	"net/http"
)

type MuxRouterAdapter struct {
	router *mux.Router
}

func (m *MuxRouterAdapter) ExtractURLParam(ctx context.Context, key string) string {
	return mux.URLParam(ctx, key)
}

func NewMuxRouterAdapter() *MuxRouterAdapter {
	router := mux.NewRouter()
	return &MuxRouterAdapter{router: router}
}

func (m *MuxRouterAdapter) Handle(pattern string, handlers MethodHandlers) {
	muxMethodHandlers := (mux.MethodHandler)(handlers)
	m.router.Handle(pattern, muxMethodHandlers)
}

func (m *MuxRouterAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}
