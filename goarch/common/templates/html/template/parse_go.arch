package template

import (
	"html/template"
	"io/fs"
)

func Parse(name string, templateFuncs template.FuncMap, templateFS fs.FS, includedPattern string, templatePattern string) *template.Template {
	tmpl := template.New(name).Funcs(templateFuncs)
	tmpl = template.Must(tmpl.ParseFS(templateFS, includedPattern))
	return template.Must(tmpl.ParseFS(templateFS, templatePattern+name))
}
