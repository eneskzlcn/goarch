package microservice

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/common"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: main map that contains the relation of directory name --> file name --> file
//go:embed templates/domain/handler_go.arch
var handlerGoFileContent string

//go:embed templates/domain/service_go.arch
var serviceGoFileContent string

//go:embed templates/domain/repository_go.arch
var repositoryGoFileContent string

var internalDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"config": common.ConfigDirectory,
		"util":   common.UtilDirectory,
		"mocks":  common.MocksDirectory,
		"client": common.ClientDirectory,
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
