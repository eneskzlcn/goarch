package microservice

import (
	"fmt"
	"github.com/eneskzlcn/goarch/inner/cmd"
	continous_delivery "github.com/eneskzlcn/goarch/inner/continous-delivery"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/dev"
	"github.com/eneskzlcn/goarch/inner/internal"
	"github.com/eneskzlcn/goarch/inner/logger"
	"github.com/eneskzlcn/goarch/inner/postgres"
	"github.com/eneskzlcn/goarch/inner/rabbitmq"
	"github.com/eneskzlcn/goarch/inner/seed"
	"github.com/eneskzlcn/goarch/inner/server"
	"github.com/eneskzlcn/goarch/inner/tech"
)

func CreateArchitecture(options tech.Options) error {
	if err := dev.PrepareDirectory(arch.Microservice); err != nil {
		fmt.Println(err)
	}
	if err := cmd.PrepareDirectory(); err != nil {
		return err
	}
	if err := server.PrepareDirectory(arch.Microservice, options); err != nil {
		fmt.Println(err)
	}
	if err := continous_delivery.PrepareDirectory(options); err != nil {
		fmt.Println(err)
		return err
	}
	if err := postgres.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}
	if err := logger.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}
	if err := seed.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}
	if err := rabbitmq.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}

	if err := internal.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}

	return nil
}
