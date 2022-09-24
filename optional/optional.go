package optional

import (
	"fmt"
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/common"
	"github.com/eneskzlcn/goarch/tech"
)

type ArchDirectoryMapper = func(p arch.Type) arch.Directory

var optionDirectoryMap = map[tech.Option]ArchDirectoryMapper{
	tech.Fiber:    common.FiberDirectory,
	tech.Gin:      common.GinDirectory,
	tech.Rabbitmq: common.RabbitmqDirectory,
	tech.Kafka:    common.KafkaDirectory,
	tech.Logger:   common.LoggerDirectory,
	//tech.Client:   common.ClientDirectory,
	tech.Postgres: common.PostgresDirectory,
}

func GetDirectoriesByOptionAndArch(t arch.Type, options tech.Options) arch.Directories {
	directories := make(arch.Directories, 0)
	for _, option := range options {
		if archDirMapper, exists := optionDirectoryMap[option]; exists {
			directory := archDirMapper(t)
			directories = append(directories, directory)
		}
	}
	for _, item := range directories {
		fmt.Printf("Optinal Directory add to %s with name %s\n", item.AbsPath, item.Name)
	}
	return directories
}
