package microservice

import (
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/common"
	"github.com/eneskzlcn/goarch/tech"
)

const architecture = arch.Microservice

func CreateArchitecture(options tech.Options) error {
	a := arch.New()
	a.Add(common.RootDirectories...)

	a.Add(InternalDirectory)
	//optionalDirectories := optional.GetDirectoriesByOptionAndArch(architecture, options)
	//a.Add(optionalDirectories...)

	return a.Create()
}
