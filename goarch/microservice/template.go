package microservice

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/common"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: main map that contains the relation of directory name --> file name --> file
//go:embed templates/domain/handler.go
var handlerGoFileContent string

//go:embed templates/domain/service.go
var serviceGoFileContent string

//go:embed templates/domain/repository.go
var repositoryGoFileContent string

var internalDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"config": common.ConfigDirectory,
		"util":   common.UtilDirectory,
		"mocks":  common.MocksDirectory,
		"domain": directory.Directory{
			Files: file.Files{
				"domain":          file.NewGoFile(""),
				"handler":         file.NewGoFile(handlerGoFileContent),
				"handler_test":    file.NewGoTestFile(""),
				"service":         file.NewGoFile(serviceGoFileContent),
				"service_test":    file.NewGoTestFile(""),
				"repository":      file.NewGoFile(repositoryGoFileContent),
				"repository_test": file.NewGoFile(""),
			},
		},
	},
}
