package nlayered_common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: repository directory initialization.

//go:embed templates/service/service_go.arch
var serviceGoFileContent string

var ServiceDirectory = directory.Directory{
	Files: file.Files{
		"repository":      file.NewGoFile(serviceGoFileContent),
		"repository_test": file.NewGoTestFile(""),
	},
}
