package nlayered_web

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/common"
	nlayered_common "github.com/eneskzlcn/goarch/goarch/common/nlayered-common"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

var coreDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"logger":   common.LoggerDirectory,
		"postgres": common.PostgresDirectory,
		"util":     common.UtilDirectory,
		"server":   common.ServerDirectory,
		"rabbitmq": common.ConfigDirectory,
		"client":   common.ClientDirectory,
		"session":  common.SessionDirectory,
		"router":   common.RouterDirectory,
		"cache":    common.CacheDirectory,
		"mail":     common.MailDirectory,
	},
}

var internalDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"client":     common.ClientDirectory,
		"config":     common.ConfigDirectory,
		"core":       coreDirectory,
		"service":    nlayered_common.ServiceDirectory,
		"repository": nlayered_common.RepositoryDirectory,
		"entity":     nlayered_common.EntityDirectory,
		"web":        webDirectory,
	},
}

//go:embed templates/web/template/include/README.md
var webTemplateIncludeMarkdownFileContent string

//go:embed templates/web/template/include/header.gohtml
var webTemplateIncludeHeaderFileContent string

//go:embed templates/web/template/include/layout.gohtml
var webTemplateIncludeLayoutFileContent string

//go:embed templates/web/template/home.gohtml
var webTemplateHomeFileContent string

var templateDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"include": directory.Directory{
			Files: file.Files{
				"layout": file.NewGoHtmlFile(webTemplateIncludeLayoutFileContent),
				"header": file.NewGoHtmlFile(webTemplateIncludeHeaderFileContent),
				"README": file.NewMarkdownFile(webTemplateIncludeMarkdownFileContent),
			},
		},
	},
	Files: file.Files{
		"home": file.NewGoHtmlFile(webTemplateHomeFileContent),
	},
}

//MARK: web directory initialization.

//go:embed templates/web/template_go.arch
var webTemplateGoFileContent string

//go:embed templates/web/handler_go.arch
var webHandlerGoFileContent string

//go:embed templates/web/home_go.arch
var webHomeGoFileContent string

//go:embed templates/web/web_go.arch
var webGoFileContent string

var webDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"template": templateDirectory,
	},
	Files: file.Files{
		"template": file.NewGoFile(webTemplateGoFileContent),
		"handler":  file.NewGoFile(webHandlerGoFileContent),
		"home":     file.NewGoFile(webHomeGoFileContent),
		"web":      file.NewGoFile(webGoFileContent),
	},
}
