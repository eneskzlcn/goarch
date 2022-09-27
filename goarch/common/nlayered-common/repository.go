package nlayered_common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: repository directory initialization.

//go:embed templates/repository/repository_go.arch
var repositoryGoFileContent string

var RepositoryDirectory = directory.Directory{
	Files: file.Files{
		"repository":      file.NewGoFile(repositoryGoFileContent),
		"repository_test": file.NewGoTestFile(""),
	},
}
