package server

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/internal/arch"
	"github.com/eneskzlcn/goarch/internal/file"
	"github.com/eneskzlcn/goarch/internal/tech"
)

const (
	directoryName      = "server"
	serverFileName     = "server.go"
	serverTestFileName = "server_test.go"
)

//go:embed fiber_server.arch
var fiberServerContent string

//go:embed gin_server.arch
var ginServerContent string

//go:embed server_test.arch
var serverTestContent string

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice: ".",
}

func PrepareDirectory(architecture arch.Type, options tech.Options) error {
	if err := file.CreateDirectory(architectureAbsPathMap[architecture], directoryName); err != nil {
		return err
	}
	if options.Has(tech.Fiber) {
		err := file.CreateFileWithContent(architectureAbsPathMap[architecture]+"/"+directoryName,
			serverFileName, fiberServerContent)
		if err != nil {
			return err
		}
	} else if options.Has(tech.Gin) {
		err := file.CreateFileWithContent(architectureAbsPathMap[architecture]+"/"+directoryName,
			serverFileName, fiberServerContent)
		if err != nil {
			return err
		}
	}
	return file.CreateFileWithContent(architectureAbsPathMap[architecture]+"/"+directoryName,
		serverTestFileName, serverTestContent)
}
