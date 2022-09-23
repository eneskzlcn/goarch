package client

import (
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

const (
	directoryName         = "client"
	kafkaDirectoryName    = "kafka"
	kafkaFileName         = "kafka.go"
	kafkaFileContent      = `package kafka`
	rabbitmqDirectoryName = "rabbitmq"
	rabbitmqFileName      = "rabbitmq.go"
	rabbitmqFileContent   = `package rabbitmq`
	restDirectoryClient   = "rest"
	restFileName          = "rest.go"
	restFileContent       = `package rest`
)

var clientNames = []string{"kafka", "rabbitmq", "rest"}

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   "./internal",
	arch.NLayeredWebApp: "./internal/core",
}

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}
	clientDir := baseDir + "/" + directoryName
	for _, name := range clientNames {
		if err := utl.CreateDirectory(clientDir, name); err != nil {
			return err
		}
	}
	for _, name := range clientNames {
		targetDir := clientDir + "/" + name
		if err := utl.CreateFileWithContent(targetDir, name+".go", "package "+name); err != nil {
			return err
		}
	}
	return nil
}
