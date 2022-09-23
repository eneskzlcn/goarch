package dev

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/command"
)

//go:embed default_config_yaml.architecture
var defaultConfigYamlContent string

var directoryName = ".dev"

var configFileNames = []string{"dev.yaml", "qa.yaml", "local.yaml", "prod.yaml"}

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice: ".",
}

func PrepareDirectory(architecture arch.Type) error {
	if err := command.CreateDirectory(architectureAbsPathMap[architecture], directoryName); err != nil {
		return err
	}
	for _, configFileName := range configFileNames {
		if err := command.CreateFileWithContent(architectureAbsPathMap[architecture]+"/"+directoryName,
			configFileName, defaultConfigYamlContent); err != nil {
			return err
		}
	}
	return nil
}

//
//func PrepareDevDirectory() error {
//	if err := CreateDirectory(".dev"); err != nil {
//		return err
//	}
//
//	if err := CreateFileWithContent(".dev/dev.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := CreateFileWithContent(".dev/qa.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := CreateFileWithContent(".dev/prod.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := CreateFileWithContent(".dev/local.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	return nil
//}
