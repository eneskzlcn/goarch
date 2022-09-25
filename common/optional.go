package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/file"
	"github.com/eneskzlcn/goarch/goarch/directory"
)

const (
	GinDirectoryName      = "server"
	FiberDirectoryName    = GinDirectoryName
	PostgresDirectoryName = "postgres"
	KafkaDirectoryName    = "kafka"
	RabbitmqDirectoryName = "rabbitmq"
)

//MARK: ./kafka directory initialization.

func KafkaDirectory(architectureType architecture.Type) architecture.Directory {
	emptyGoFileNames := []string{KafkaDirectoryName}
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(KafkaDirectoryName, architectureType),
		Name:    KafkaDirectoryName,
		SubDir:  nil,
		Files:   file.CreateEmptyGoAndTestFilesByGivenFileNames(emptyGoFileNames...),
	}

}

//MARK: ./rabbitmq directory initialization.

//go:embed templates/rabbitmq/rabbitmq_go.arch
var rabbitmqFileContent string

func RabbitmqDirectory(architectureType architecture.Type) architecture.Directory {
	emptyGoFileNames := []string{RabbitmqDirectoryName}
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(RabbitmqDirectoryName, architectureType),
		Name:    RabbitmqDirectoryName,
		SubDir:  nil,
		Files:   file.CreateEmptyGoAndTestFilesByGivenFileNames(emptyGoFileNames...),
	}
}

//MARK: ./server directory initialization for fiber.

//go:embed templates/server/fiber_server.arch
var fiberServerFileContent string

func FiberDirectory(architectureType architecture.Type) architecture.Directory {
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(FiberDirectoryName, architectureType),
		Name:    FiberDirectoryName,
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    FiberDirectoryName + ".go",
				Content: fiberServerFileContent,
			},
			{
				Name:    FiberDirectoryName + "_test.go",
				Content: "package " + FiberDirectoryName + "_test",
			},
		},
	}
}

//go:embed templates/server/gin_server.arch
var ginServerFileContent string

func GinDirectory(architectureType architecture.Type) architecture.Directory {
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(GinDirectoryName, architectureType),
		Name:    "server",
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    "server.go",
				Content: ginServerFileContent,
			},
			{
				Name:    "server_test.go",
				Content: "package server_test",
			},
		},
	}
}

//MARK: ./postgres directory initialization.

//go:embed templates/postgres/postgres_go.arch
var postgresGoFileContent string

//go:embed templates/postgres/postgres_test.arch
var postgresTestFileContent string

//go:embed templates/postgres/mock_postgres_go.arch
var mockPostgresFileContent string

func PostgresDirectory(architectureType architecture.Type) architecture.Directory {
	return architecture.Directory{
		AbsPath: directory.FindPathOfGivenDirectoryByNameAndArchitecture(PostgresDirectoryName, architectureType),
		Name:    PostgresDirectoryName,
		SubDir:  nil,
		Files: architecture.Files{
			{
				Name:    PostgresDirectoryName + ".go",
				Content: postgresGoFileContent,
			},
			{
				Name:    PostgresDirectoryName + "_test.go",
				Content: postgresTestFileContent,
			},
			{
				Name:    "mock_" + PostgresDirectoryName + ".go",
				Content: mockPostgresFileContent,
			},
		},
	}
}
