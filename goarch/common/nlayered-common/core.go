package nlayered_common

import (
	"github.com/eneskzlcn/goarch/goarch/common"
	"github.com/eneskzlcn/goarch/goarch/directory"
)

// MARK: core directory initialization

var CoreDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"logger":   common.LoggerDirectory,
		"postgres": common.PostgresDirectory,
		"util":     common.UtilDirectory,
		"server":   common.ServerDirectory,
		"rabbitmq": common.ConfigDirectory,
		"client":   common.ClientDirectory,
	},
	Files: nil,
}
