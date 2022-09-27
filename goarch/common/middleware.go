package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: middleware directory initialization

//go:embed templates/middleware/override_form_methods_go.arch
var middlewareOverrideFormMethodsGoFileContent string

var MiddlewareDirectory = directory.Directory{
	Files: file.Files{
		"override_form_methods": file.NewGoFile(middlewareOverrideFormMethodsGoFileContent),
	},
}
