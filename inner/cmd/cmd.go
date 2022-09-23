package cmd

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/utl"
	"os"
	"path"
)

//go:embed main_go.arch
var mainFileContent string

const (
	defaultDirectory = "."
	directoryName    = "cmd"
	cmdDirectory     = defaultDirectory + "/" + directoryName
	mainFileName     = "main.go"
)

func PrepareDirectory() error {
	if err := utl.CreateDirectory(defaultDirectory, directoryName); err != nil {
		return err
	}
	currentDirectory, _ := os.Getwd()
	projectDirName := path.Base(currentDirectory)

	if err := utl.CreateDirectory(cmdDirectory, projectDirName); err != nil {
		return err
	}
	mainFileDirectory := cmdDirectory + "/" + projectDirName

	if err := utl.CreateFileWithContent(mainFileDirectory, mainFileName, mainFileContent); err != nil {
		return err
	}
	return nil
}
