package microservice

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/file"
)

type DirectoryNameFileMap map[string]NameFileMap

type NameFileMap map[string]file.File

//MARK: dev directory file content map initialization.

//go:embed templates/.dev/default_config_yaml.arch
var defaultDevConfigYamlContent string

var DevDirectoryNameFileMap = map[string]file.File{
	"local": file.NewYamlFile(defaultDevConfigYamlContent),
	"qa":    file.NewYamlFile(defaultDevConfigYamlContent),
	"prod":  file.NewYamlFile(defaultDevConfigYamlContent),
	"dev":   file.NewYamlFile(defaultDevConfigYamlContent),
}

//MARK: logger directory file contetn map initialization

//var LoggerFileContentMap = map[string]string {
//	"logger":
//}

var directoryToNameFileMapper = DirectoryNameFileMap{
	".dev": DevDirectoryNameFileMap,
}
