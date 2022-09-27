package web

import (
	"net/http"
)

type homeData struct {
	Header string
}

func (h *Handler) ShowHome(w http.ResponseWriter, r *http.Request) {
	h.RenderHome(w, homeData{Header: "Welcome"}, http.StatusFound)
}

func (h *Handler) RenderHome(w http.ResponseWriter, data homeData, status int) {
	h.RenderPage("home", w, data, status)
}
