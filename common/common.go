package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/arch"
)

//MARK: ./logger or ./internal/core/logger directory initialization.

//go:embed templates/logger/logger_go.arch
var loggerFileContent string

//go:embed templates/logger/zap_logger_go.arch
var loggerZapFileContent string

//go:embed templates/logger/zap_logger_test.arch
var loggerZapTestFileContent string

func LoggerDirectory(t arch.Type) arch.Directory {
	dirName := "logger"
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[t][dirName],
		Name:    dirName,
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    dirName + ".go",
				Content: loggerFileContent,
			},
			{
				Name:    "zap_logger_adapter.go",
				Content: loggerZapFileContent,
			},
			{
				Name:    "zap_logger_adapter_test.go",
				Content: loggerZapTestFileContent,
			},
		},
	}
}

//MARK: ./internal/client or ./internal/core/client directory initialization.

func ClientDirectory(t arch.Type) arch.Directory {
	dirName := "client"
	kafkaDirName := "kafkaclient"
	httpDirName := "httpclient"
	rabbitmqDirName := "rabbitmqclient"
	clientDirPath := arch.PathByArchAndDirName[t][dirName] + "/" + dirName
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[t][dirName],
		Name:    dirName,
		SubDir:  directoriesWithEmptyGoAndTestFiles(clientDirPath, kafkaDirName, httpDirName, rabbitmqDirName),
	}
}

//MARK: ./internal/util or ./internal/core/util directory initialization.

func UtilDirectory(t arch.Type) arch.Directory {
	dirName := "util"
	dirPath := arch.PathByArchAndDirName[t][dirName]
	utilDirPath := dirPath + "/" + dirName
	return arch.Directory{
		AbsPath: dirPath,
		Name:    dirName,
		SubDir:  directoriesWithEmptyGoAndTestFiles(utilDirPath, "ctxutil"),
	}
}

func directoriesWithEmptyGoAndTestFiles(dirPath string, dirNames ...string) arch.Directories {
	directories := make(arch.Directories, 0)
	for _, name := range dirNames {
		directories = append(directories, arch.Directory{
			AbsPath: dirPath + "/",
			Name:    name,
			SubDir:  nil,
			Files: arch.Files{
				{
					Name:    name + ".go",
					Content: "package " + name,
				},
				{
					Name:    name + "_test.go",
					Content: "package " + name + "_test",
				},
			},
		})
	}
	return directories
}
