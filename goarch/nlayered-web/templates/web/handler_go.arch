package web

import (
	logger "<path-to-core>/logger"
	middleware "<path-to-core>/middleware"
	router "<path-to-core>/router"
	"context"
	"encoding/gob"
	"errors"
	"html/template"
	"net/http"
	"net/url"
)

type Service interface {
	GetUsers(ctx context.Context, limit int)
}
type Renderer interface {
	Render(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int)
}

type Handler struct {
	handler           http.Handler
	logger            logger.Logger
	renderer          Renderer
	service           Service
	templates         PageTemplates
	urlParamExtractor func(ctx context.Context, key string) string
}

func NewHandler(logger logger.Logger, service Service, renderer Renderer) *Handler {
	if logger == nil {
		return nil
	}
	if session == nil || service == nil || renderer == nil {
		logger.Error("invalid arguments"))
		return nil
	}
	handler := Handler{logger: logger, service: service, renderer: renderer}

	handler.init()

	return &handler
}
func (h *Handler) init() {
	muxRouter := router.NewMuxRouterAdapter()
	h.urlParamExtractor = muxRouter.ExtractURLParam
	h.handler = muxRouter
	h.RegisterHandlers(muxRouter)

	h.applyMiddlewares()

	gob.Register(url.Values{})

	templates := make(PageTemplates, 0)
	templates["home"] = ParseTemplate("home.gohtml")
	h.templates = templates
}
func (h *Handler) applyMiddlewares() {
	//apply method overriding middleware
	h.handler = middleware.OverrideFormMethods(h.handler)
}

func (h *Handler) RenderPage(page string, w http.ResponseWriter, data any, statusCode int) {
	h.renderer.Render(w, h.templates[page], data, statusCode)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handler.ServeHTTP(w, req)
}

func (h *Handler) RegisterHandlers(muxRouter router.Router) {
	muxRouter.Handle("/", router.MethodHandlers{
		http.MethodGet: h.ShowHome,
	})

}
