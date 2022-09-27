package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: cache directory initialization

//go:embed templates/html/template/linkify_go.arch
var htmlTemplateLinkifyGoFileContent string

//go:embed templates/html/template/parse_go.arch
var htmlTemplateParseGoFileContent string

//go:embed templates/html/template/renderer/renderer_go.arch
var htmlTemplateRendererGoFileContent string

var htmlTemplateRendererDirectory = directory.Directory{
	Files: file.Files{
		"renderer": file.NewGoFile(htmlTemplateRendererGoFileContent),
	},
}

var htmlTemplateDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"renderer": htmlTemplateRendererDirectory,
	},
	Files: file.Files{
		"linkify": file.NewGoFile(htmlTemplateLinkifyGoFileContent),
		"parse":   file.NewGoFile(htmlTemplateParseGoFileContent),
	},
}
var HtmlDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"template": htmlTemplateDirectory,
	},
}
