package microservice

import (
	"fmt"
	"github.com/eneskzlcn/goarch/inner/arch"
	"github.com/eneskzlcn/goarch/inner/cmd"
	continous_delivery "github.com/eneskzlcn/goarch/inner/continous-delivery"
	"github.com/eneskzlcn/goarch/inner/dev"
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
	//if err := PrepareInternalDirectory(); err != nil {
	//	return err
	//}
	//
	//if err := PreparePostgresDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareLoggerDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareRabbitMQDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareSeedDirectory(); err != nil {
	//	return err
	//}
	return nil
}
