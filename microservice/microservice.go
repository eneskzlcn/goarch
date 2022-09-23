package microservice

import (
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/dev"
	"github.com/eneskzlcn/goarch/tech"
)

func CreateArchitecture(technologyOptions ...tech.Options) error {

	if err := dev.PrepareDirectory(arch.Microservice); err != nil {
		return err
	}
	//if err := CreateCmdDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareServerDirectory(Gin); err != nil {
	//	return err
	//}

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
