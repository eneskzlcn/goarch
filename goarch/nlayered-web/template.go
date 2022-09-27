package nlayered_web

import (
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

var templateDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"include": directory.Directory{
			Files: file.Files{
				"layout": file.NewGoHtmlFile(""),
				"header": file.NewGoHtmlFile(""),
			},
		},
	},
	Files: file.Files{
		"home": file.NewGoHtmlFile(""),
	},
}
var webDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"template": templateDirectory,
	},
	Files: file.Files{
		//TODO: content will be taken by embed,
		"template": file.NewGoFile(""),
		"handler":  file.NewGoFile(""),
		"home":     file.NewGoFile(""),
	},
}
