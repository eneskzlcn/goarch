package internal

import (
	"github.com/eneskzlcn/goarch/inner/client"
	"github.com/eneskzlcn/goarch/inner/config"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
	"github.com/eneskzlcn/goarch/inner/domain"
	"github.com/eneskzlcn/goarch/inner/mocks"
	"github.com/eneskzlcn/goarch/inner/util"
)

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   ".",
	arch.NLayeredWebApp: ".",
}

const (
	directoryName = "internal"
)

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}
	if err := config.PrepareDirectory(architecture); err != nil {
		return err
	}

	if err := domain.PrepareDirectory(); err != nil {
		return err
	}

	if err := mocks.PrepareDirectory(); err != nil {
		return err
	}
	if err := util.PrepareDirectory(architecture); err != nil {
		return err
	}

	if err := client.PrepareDirectory(architecture); err != nil {
		return err
	}

	return nil
}
