package common

import (
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

var ClientDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"rabbitclient": directory.Directory{
			Files: file.Files{
				"rabbitclient":      file.NewGoFile(""),
				"rabbitclient_test": file.NewGoTestFile(""),
			},
		},
		"httpclient": directory.Directory{
			Files: file.Files{
				"httpclient":      file.NewGoFile(""),
				"httpclient_test": file.NewGoTestFile(""),
			},
		},
	},
}
