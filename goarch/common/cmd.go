package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: cmd directory initialization.

//go:embed templates/cmd/main_go.arch
var cmdMainGoFileContent string

var CmdDirectory = directory.Directory{
	Files: file.Files{
		"main": file.NewGoFile(cmdMainGoFileContent),
	},
}
