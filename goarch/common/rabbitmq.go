package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: rabbitmq directory initialization.

//go:embed templates/rabbitmq/rabbitmq_go.arch
var rabbitmqGoFileContent string

var RabbitmqDirectory = directory.Directory{
	Files: file.Files{
		"rabbitmq": file.NewGoFile(rabbitmqGoFileContent),
	},
}
