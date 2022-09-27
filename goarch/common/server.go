package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK server directory initialization

//go:embed templates/server/fiber_server.arch
var fiberServerGoFileContent string

var ServerDirectory = directory.Directory{
	Files: file.Files{
		"server": file.NewGoFile(fiberServerGoFileContent),
	},
}
