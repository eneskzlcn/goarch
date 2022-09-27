package domain

import "github.com/gofiber/fiber/v2"

type DomainService interface {
	Example() (string, error)
}
type Logger interface {
}
type Handler struct {
	service DomainService
	logger  Logger
}

func NewHandler(service DomainService, logger Logger) *Handler {
	if logger == nil || service == nil {
		return nil
	}
	return &Handler{
		service: service,
		logger:  logger,
	}
}
func (h *Handler) Example(ctx *fiber.Ctx) error {
	example, err := h.service.Example()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendString(example)
}
func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/example", h.Example)
}
