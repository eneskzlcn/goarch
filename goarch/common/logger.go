package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: logger directory initialization.

//go:embed templates/logger/logger_go.arch
var loggerLoggerGoFileContent string

//go:embed templates/logger/zap_logger_go.arch
var loggerZapLoggerAdapterFileContent string

//go:embed templates/logger/zap_logger_test.arch
var loggerZapLoggerAdapterTestFileContent string

var LoggerDirectory = directory.Directory{
	Files: file.Files{
		"logger.go":                  file.NewGoFile(loggerLoggerGoFileContent),
		"zap_logger_adapter_test.go": file.NewGoTestFile(loggerZapLoggerAdapterTestFileContent),
		"zap_logger_adapter.go":      file.NewGoFile(loggerZapLoggerAdapterFileContent),
	},
}
