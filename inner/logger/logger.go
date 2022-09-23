package logger

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed logger_go.arch
var loggerFileContent string

//go:embed zap_logger_go.arch
var zapLoggerFileContent string

//go:embed zap_logger_test.arch
var zapLoggerTestFileContent string

const (
	directoryName         = "logger"
	loggerFileName        = "logger.go"
	zapLoggerFileName     = "zap_logger_adapter.go"
	zapLoggerTestFileName = "zap_logger_adapter_test.go"
)

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   ".",
	arch.NLayeredWebApp: "./internal/core",
}

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}

	loggerDir := baseDir + "/" + directoryName
	if err := utl.CreateFileWithContent(loggerDir, loggerFileName, loggerFileContent); err != nil {
		return err
	}

	if err := utl.CreateFileWithContent(loggerDir, zapLoggerFileName, zapLoggerFileContent); err != nil {
		return err
	}

	if err := utl.CreateFileWithContent(loggerDir, zapLoggerTestFileName, zapLoggerTestFileContent); err != nil {
		return err
	}

	return nil
}
