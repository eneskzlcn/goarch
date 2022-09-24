package anotherway

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/tech"
)

//go:embed directory.go
var content string

var optionDirectoryMap = map[tech.Option]Directory{
	tech.Fiber:    FiberServerDirectory,
	tech.Gin:      GinServerDirectory,
	tech.Rabbitmq: RabbitmqDirectory,
	tech.Kafka:    KafkaDirectory,
}

func CreateMicroserviceArchitecture(options tech.Options) error {
	arch := Architecture{Directories{
		DevDirectory,
		SeedDirectory,
		InternalDirectory,
		LoggerDirectory,
	}}
	optionalDirectories := GetOptionalDirectories(options)
	arch.Directories = append(arch.Directories, optionalDirectories...)
	return arch.Create()
}
func GetOptionalDirectories(options tech.Options) []Directory {
	directories := make([]Directory, 0)
	for _, option := range options {
		if dir, exists := optionDirectoryMap[option]; exists {
			directories = append(directories, dir)
		}
	}
	return directories
}

var KafkaDirectory = Directory{
	AbsPath: "",
	Name:    "",
	SubDir:  nil,
	Files:   nil,
}
var InternalDirectory = Directory{
	AbsPath: "",
	Name:    "",
	SubDir:  nil,
	Files:   nil,
}
var SeedDirectory = Directory{
	AbsPath: ".",
	Name:    "seed",
	SubDir: Directories{
		{
			AbsPath: "./seed",
			Name:    "cmd",
			SubDir:  nil,
			Files: Files{
				{
					Name:    "main.go",
					Content: "package main",
				},
			},
		},
	},
	Files: Files{
		{
			Name:    "create-seed.sql",
			Content: content,
		},
		{
			Name:    "drop-seed.sql",
			Content: content,
		},
		{
			Name:    "seed.go",
			Content: "package seed",
		},
	},
}
var RabbitmqDirectory = Directory{
	AbsPath: ".",
	Name:    "rabbitmq",
	SubDir:  nil,
	Files: Files{
		{
			Name:    "rabbitmq.go",
			Content: "package rabbitmq",
		},
		{
			Name:    "rabbitmq_test.go",
			Content: "package rabbitmq_test",
		},
	},
}
var GinServerDirectory = Directory{
	AbsPath: ".",
	Name:    "server",
	SubDir:  nil,
	Files: Files{
		{
			Name:    "server.go",
			Content: "package server",
		},
		{
			Name:    "server_test.go",
			Content: "package server_test",
		},
	},
}
var FiberServerDirectory = Directory{
	AbsPath: ".",
	Name:    "server",
	SubDir:  nil,
	Files: Files{
		{
			Name:    "server.go",
			Content: "package server",
		},
		{
			Name:    "server_test.go",
			Content: "package server_test",
		},
	},
}
var LoggerDirectory = Directory{
	AbsPath: ".",
	Name:    "logger",
	SubDir:  nil,
	Files: Files{
		{
			Name:    "logger.go",
			Content: "package logger",
		},
	},
}
var DevDirectory = Directory{
	AbsPath: ".",
	Name:    ".dev",
	SubDir:  nil,
	Files: Files{
		{
			Name:    "dev.yaml",
			Content: content,
		},
		{
			Name:    "qa.yaml",
			Content: content,
		},
		{
			Name:    "prod.yaml",
			Content: content,
		},
		{
			Name:    "test.yaml",
			Content: content,
		},
	},
}

//var Microservice = Architecture{Directories: []Directory{
//	{
//		AbsPath: ".",
//		Name:    ".dev",
//		SubDir:  nil,
//	},
//}}
