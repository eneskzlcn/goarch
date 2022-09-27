package nlayered_common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: entity directory initialization.

//go:embed templates/entity/user_go.arch
var entityUserGoFileContent string

var EntityDirectory = directory.Directory{
	Files: file.Files{
		"user": file.NewGoFile(entityUserGoFileContent),
	},
}
