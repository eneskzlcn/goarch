package optional

import (
	"fmt"
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/common"
	"github.com/eneskzlcn/goarch/tech"
)

type DirectoryCreatorByArchitecture = func(p architecture.Type) architecture.Directory

var optionDirectoryMap = map[tech.Option]DirectoryCreatorByArchitecture{
	tech.Fiber:    common.FiberDirectory,
	tech.Gin:      common.GinDirectory,
	tech.Rabbitmq: common.RabbitmqDirectory,
	tech.Kafka:    common.KafkaDirectory,
	tech.Logger:   common.LoggerDirectory,
	tech.Postgres: common.PostgresDirectory,
}

func GetDirectoriesByOptionAndArch(t architecture.Type, options tech.Options) architecture.Directories {
	directories := make(architecture.Directories, 0)
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
