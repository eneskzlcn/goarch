package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/goarch/directory"
)

const (
	ConfigDirectoryName = "config"
	MocksDirectoryName  = "mocks"
)

//MARK: ./internal/config directory initialization.

//go:embed templates/config/config_go.arch
var configFileContent string

//go:embed templates/config/config_test.arch
var configTestFileContent string

//go:embed templates/config/db_go.arch
var configDBFileContent string

func ConfigDirectory(architectureType architecture.Type) architecture.Directory {
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(ConfigDirectoryName, architectureType),
		Name:    ConfigDirectoryName,
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    ConfigDirectoryName + ".go",
				Content: configFileContent,
			},
			{
				Name:    ConfigDirectoryName + "_test.go",
				Content: configTestFileContent,
			},
			{
				Name:    "db.go",
				Content: configDBFileContent,
			},
		},
	}
}

//MARK: ./internal/mocks directory initialization.

func MocksDirectory(architectureType architecture.Type) architecture.Directory {
	dirPath := directory.FindPathOfGivenDirectoryByNameAndArchitecture(MocksDirectoryName, architectureType)
	exFile := "mock_user_repository"
	return architecture.Directory{
		AbsPath: dirPath,
		Name:    MocksDirectoryName,
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    exFile + ".go",
				Content: "package " + MocksDirectoryName,
			},
			{
				Name:    exFile + "_test.go",
				Content: "package " + MocksDirectoryName + "_test",
			},
		},
	}
}
