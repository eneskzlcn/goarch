package goarch

import (
	"errors"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/microservice"
	"github.com/eneskzlcn/goarch/inner/core/nlayered"
	"github.com/eneskzlcn/goarch/inner/core/tech"
)

type ArchitectureCreator func(options tech.Options) error

var architectureCreatorMap = map[arch.Type]ArchitectureCreator{
	arch.Microservice:    microservice.CreateArchitecture,
	arch.NLayeredWebApp:  nlayered.CreateWebArchitecture,
	arch.NLayeredBackend: nlayered.CreateBackendArchitecture,
}

func Create(architecture arch.Type, options tech.Options) error {
	if creatorFn, exists := architectureCreatorMap[architecture]; exists {
		return creatorFn(options)
	}
	return errors.New("not valid operation")
}
