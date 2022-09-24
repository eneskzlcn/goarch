package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/arch"
)

//MARK: ./kafka directory initialization.

func KafkaDirectory(t arch.Type) arch.Directory {
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[t]["kafka"],
		Name:    "kafka",
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    "kafka.go",
				Content: "package kafka",
			},
			{
				Name:    "kafka.go",
				Content: "package kafka_test",
			},
		},
	}
}

//MARK: ./rabbitmq directory initialization.

//go:embed templates/rabbitmq/rabbitmq_go.arch
var rabbitmqFileContent string

func RabbitmqDirectory(t arch.Type) arch.Directory {
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[t]["rabbitmq"],
		Name:    "rabbitmq",
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    "rabbitmq.go",
				Content: rabbitmqFileContent,
			},
			{
				Name:    "rabbitmq_test.go",
				Content: "package rabbitmq_test",
			},
		},
	}
}

//MARK: ./server directory initialization for fiber.

//go:embed templates/server/fiber_server.arch
var fiberServerFileContent string

func FiberDirectory(p arch.Type) arch.Directory {
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[p]["server"],
		Name:    "server",
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    "server.go",
				Content: fiberServerFileContent,
			},
			{
				Name:    "server_test.go",
				Content: "package server_test",
			},
		},
	}
}

//go:embed templates/server/gin_server.arch
var ginServerFileContent string

func GinDirectory(p arch.Type) arch.Directory {
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[p]["server"],
		Name:    "server",
		SubDir:  nil,
		Files: arch.Files{
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

func PostgresDirectory(p arch.Type) arch.Directory {
	dirName := "postgres"
	return arch.Directory{
		AbsPath: arch.PathByArchAndDirName[p][dirName],
		Name:    dirName,
		SubDir:  nil,
		Files: arch.Files{
			{
				Name:    dirName + ".go",
				Content: postgresGoFileContent,
			},
			{
				Name:    dirName + "_test.go",
				Content: postgresTestFileContent,
			},
			{
				Name:    "mock_" + dirName + ".go",
				Content: mockPostgresFileContent,
			},
		},
	}
}
