package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/arch"
)

//MARK: ./internal/config directory initialization.

//go:embed templates/config/config_go.arch
var configFileContent string

//go:embed templates/config/config_test.arch
var configTestFileContent string

//go:embed templates/config/db_go.arch
var configDBFileContent string

func ConfigDirectory(t arch.Type) arch.Directory {
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[t]["config"],
		Name:    "config",
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    "config.go",
				Content: configFileContent,
			},
			{
				Name:    "config_test.go",
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

func MocksDirectory(t arch.Type) arch.Directory {
	dirName := "mocks"
	dirPath := arch.PathByArchAndDirName[t][dirName]
	exFile := "mock_user_repository"
	return arch.Directory{
		AbsPath: dirPath,
		Name:    dirName,
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    exFile + ".go",
				Content: "package " + dirName,
			},
			{
				Name:    exFile + "_test.go",
				Content: "package " + dirName + "_test",
			},
		},
	}
}
