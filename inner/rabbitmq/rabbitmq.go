package rabbitmq

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed rabbitmq_go.arch
var rabbitmqFileContent string

const (
	directoryName    = "rabbitmq"
	rabbitmqFileName = "rabbitmq.go"
)

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   ".",
	arch.NLayeredWebApp: "./internal/core",
}

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}

	rabbitmqDir := baseDir + "/" + directoryName
	if err := utl.CreateFileWithContent(rabbitmqDir, rabbitmqFileName, rabbitmqFileContent); err != nil {
		return err
	}

	return nil
}
