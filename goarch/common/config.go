package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: config directory initialization.

//go:embed templates/config/config_go.arch
var configGoFileContent string

//go:embed templates/config/config_test.arch
var configGoTestFileContent string

//go:embed templates/config/db_go.arch
var configDbGoFileContent string

var ConfigDirectory = directory.Directory{
	Files: file.Files{
		"config":      file.NewGoFile(configGoFileContent),
		"config_test": file.NewGoTestFile(configGoTestFileContent),
		"db":          file.NewGoFile(configDbGoFileContent),
	},
}
