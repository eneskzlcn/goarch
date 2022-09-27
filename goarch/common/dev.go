package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: dev directory initialization.

//go:embed templates/.dev/default_config_yaml.arch
var defaultDevConfigYamlContent string

var DevDirectory = directory.Directory{
	Files: file.Files{
		"local": file.NewYamlFile(defaultDevConfigYamlContent),
		"qa":    file.NewYamlFile(defaultDevConfigYamlContent),
		"dev":   file.NewYamlFile(defaultDevConfigYamlContent),
		"prod":  file.NewYamlFile(defaultDevConfigYamlContent),
	},
}
