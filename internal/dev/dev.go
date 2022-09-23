package dev

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/internal/arch"
	"github.com/eneskzlcn/goarch/internal/file"
)

//go:embed default_config_yaml.arch
var defaultConfigYamlContent string

var directoryName = ".dev"

var configFileNames = []string{"dev.yaml", "qa.yaml", "local.yaml", "prod.yaml"}

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice: ".",
}

func PrepareDirectory(architecture arch.Type) error {
	if err := file.CreateDirectory(architectureAbsPathMap[architecture], directoryName); err != nil {
		return err
	}
	for _, configFileName := range configFileNames {
		if err := file.CreateFileWithContent(architectureAbsPathMap[architecture]+"/"+directoryName,
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
//	if err := Options(".dev/dev.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := Options(".dev/qa.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := Options(".dev/prod.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	if err := Options(".dev/local.yaml", ConfigurationYamlContent); err != nil {
//		return err
//	}
//
//	return nil
//}
