package config

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed config_go.arch
var configFileContent string

//go:embed config_test.arch
var configTestFileContent string

//go:embed db_go.arch
var dbFileContent string

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   "./internal",
	arch.NLayeredWebApp: "./internal",
}

const (
	directoryName      = "config"
	configFileName     = "config.go"
	dbConfigFileName   = "db.go"
	configTestFileName = "config_test.go"
)

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}

	configDir := baseDir + "/" + directoryName
	if err := utl.CreateFileWithContent(configDir, configFileName, configFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(configDir, configTestFileName, configTestFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(configDir, dbConfigFileName, dbFileContent); err != nil {
		return err
	}
	return nil
}
