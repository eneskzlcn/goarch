package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: util directory initialization.

//go:embed templates/util/ctxutil/ctxutil_go.arch
var ctxUtilGoFileContent string

var UtilDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"ctxutil": directory.Directory{
			Files: file.Files{
				"ctxutil": file.NewGoFile(ctxUtilGoFileContent),
			},
		},
	},
}
