package nlayered_backend

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/common"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: core directory initialization
var coreDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"logger":   common.LoggerDirectory,
		"postgres": common.PostgresDirectory,
		"util":     common.UtilDirectory,
		"server":   common.ServerDirectory,
		"rabbitmq": common.ConfigDirectory,
		"client":   common.ClientDirectory,
	},
	Files: nil,
}

// MARK: internal directory initialization

//go:embed templates/service/service_go.arch
var serviceGoFileContent string

//go:embed templates/repository/repository_go.arch
var repositoryGoFileContent string

//go:embed templates/entity/user_go.arch
var entityUserGoFileContent string

//go:embed templates/api/api_go.arch
var apiGoFileContent string

var internalDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"client": common.ClientDirectory,
		"config": common.ConfigDirectory,
		"core":   coreDirectory,
		"service": directory.Directory{
			Files: file.Files{
				"service":      file.NewGoFile(serviceGoFileContent),
				"service_test": file.NewGoTestFile(""),
			},
		},
		"repository": directory.Directory{
			Files: file.Files{
				"repository":      file.NewGoFile(repositoryGoFileContent),
				"repository_test": file.NewGoTestFile(""),
			},
		},
		"api": directory.Directory{
			Files: file.Files{
				"api":      file.NewGoFile(apiGoFileContent),
				"api_test": file.NewGoTestFile(""),
			},
		},
		"entity": directory.Directory{
			Files: file.Files{
				"user": file.NewGoFile(entityUserGoFileContent),
			},
		},
	},
}
