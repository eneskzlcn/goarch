package microservice

import (
	"github.com/eneskzlcn/goarch/internal/arch"
	"github.com/eneskzlcn/goarch/internal/dev"
	"github.com/eneskzlcn/goarch/internal/server"
	"github.com/eneskzlcn/goarch/internal/tech"
)

func CreateArchitecture(options tech.Options) error {

	if err := dev.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}
	//if err := CreateCmdDirectory(); err != nil {
	//	return err
	//}
	if err := server.PrepareDirectory(arch.Microservice, options); err != nil {
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
