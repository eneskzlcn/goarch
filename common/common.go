package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/goarch/directory"
)

const (
	ClientDirectoryName = "client"
	UtilDirectoryName   = "util"
	LoggerDirectoryName = "logger"
)

//MARK: ./logger or ./internal/core/logger directory initialization.

//go:embed templates/logger/logger_go.arch
var loggerFileContent string

//go:embed templates/logger/zap_logger_go.arch
var loggerZapFileContent string

//go:embed templates/logger/zap_logger_test.arch
var loggerZapTestFileContent string

func LoggerDirectory(architectureType architecture.Type) architecture.Directory {
	zapLoggerFileName := "zap_logger_adapter"
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(LoggerDirectoryName, architectureType),
		Name:    LoggerDirectoryName,
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    LoggerDirectoryName + ".go",
				Content: loggerFileContent,
			},
			{
				Name:    zapLoggerFileName + ".go",
				Content: loggerZapFileContent,
			},
			{
				Name:    zapLoggerFileName + "_test.go",
				Content: loggerZapTestFileContent,
			},
		},
	}
}

//MARK: ./internal/client or ./internal/core/client directory initialization.

func ClientDirectory(architectureType architecture.Type) architecture.Directory {
	subDirectoryNames := []string{"kafkaclient", "httpclient", "rabbitmqclient"}
	directoryPath := directory.FindPathOfGivenDirectoryByNameAndArchitecture(ClientDirectoryName, architectureType)
	clientDirPath := directoryPath + "/" + ClientDirectoryName
	return architecture.Directory{
		AbsPath: directoryPath,
		Name:    ClientDirectoryName,
		SubDir:  directory.CreateDirectoriesWithEmptyGoAndTestFiles(clientDirPath, subDirectoryNames...),
	}
}

//MARK: ./internal/util or ./internal/core/util directory initialization.

func UtilDirectory(architectureType architecture.Type) architecture.Directory {
	directoryPath := directory.FindPathOfGivenDirectoryByNameAndArchitecture(UtilDirectoryName, architectureType)
	utilDirPath := directoryPath + "/" + UtilDirectoryName
	subDirectoryNames := []string{"ctxutil"}
	return architecture.Directory{
		AbsPath: directoryPath,
		Name:    UtilDirectoryName,
		SubDir:  directory.CreateDirectoriesWithEmptyGoAndTestFiles(utilDirPath, subDirectoryNames...),
	}
}
