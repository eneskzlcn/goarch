package microservice

import (
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/common"
	"github.com/eneskzlcn/goarch/optional"
	"github.com/eneskzlcn/goarch/tech"
)

func CreateArchitecture(options tech.Options) error {
	arch := architecture.New()
	arch.Add(common.RootDirectories...)

	arch.Add(InternalDirectory)

	optionalDirectories := optional.GetDirectoriesByOptionAndArch(architecture.Microservice, options)
	arch.Add(optionalDirectories...)

	return arch.Create()
}
