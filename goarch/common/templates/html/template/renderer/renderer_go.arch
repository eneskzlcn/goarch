package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Logger interface {
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

type Renderer struct {
	logger Logger
}

func New(logger Logger) *Renderer {
	return &Renderer{logger: logger}
}

func (r *Renderer) Render(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	r.logger.Debugf("Rendering the template %s", tmpl.Name())
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		r.logger.Errorf("could not render template with error: %s", err.Error())
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(statusCode)
	_, err = buf.WriteTo(w)
	if err != nil {
		r.logger.Errorf("could not send the template with error:%s", err.Error())
		return
	}
}
