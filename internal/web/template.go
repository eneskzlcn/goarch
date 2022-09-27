package web

import (
	coreTemplate "<path-to-core>/html/template"
	"embed"
	"html/template"
)

//go:embed template/include/*.gohtml template/*.gohtml

var templateFS embed.FS

var templateFuncs = template.FuncMap{
	"linkify": coreTemplate.Linkify,
}

func ParseTemplate(name string) *template.Template {
	return coreTemplate.Parse(name, templateFuncs, templateFS,
		"template/include/*.gohtml", "template/")
}
