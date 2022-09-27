package nlayered_backend

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/common"
	nlayered_common "github.com/eneskzlcn/goarch/goarch/common/nlayered-common"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: internal directory initialization

//go:embed templates/api/api_go.arch
var apiGoFileContent string

var internalDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"client":     common.ClientDirectory,
		"config":     common.ConfigDirectory,
		"core":       nlayered_common.CoreDirectory,
		"service":    nlayered_common.ServiceDirectory,
		"repository": nlayered_common.RepositoryDirectory,
		"entity":     nlayered_common.EntityDirectory,
		"api":        apiDirectory,
	},
}

var apiDirectory = directory.Directory{
	Files: file.Files{
		"api":      file.NewGoFile(apiGoFileContent),
		"api_test": file.NewGoTestFile(""),
	},
}
