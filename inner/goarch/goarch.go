package goarch

import (
	"errors"
	"github.com/eneskzlcn/goarch/inner/arch"
	"github.com/eneskzlcn/goarch/inner/microservice"
	"github.com/eneskzlcn/goarch/inner/tech"
)

func Create(architecture arch.Type, options tech.Options) error {
	switch architecture {
	case arch.NLayeredBackend:
		panic("not implemented yet")
	case arch.Microservice:
		return microservice.CreateArchitecture(options)
	case arch.NLayeredWebApp:
		panic("not implemented yet")
	}
	return errors.New("not valid operation")
}
