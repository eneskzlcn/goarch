package domain

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed general_domain_file.arch
var generalFileContent string

//go:embed general_test_file.arch
var testFileContent string

const (
	baseDirectory = "./internal"
	directoryName = "domain"
)

var fileNames = []string{"request.go", "domain.go", "service.go", "handler.go", "repository.go", "response.go"}

var testFileNames = []string{"service_test.go", "repository_test.go", "handler_test.go"}

func PrepareDirectory() error {
	if err := utl.CreateDirectory(baseDirectory, directoryName); err != nil {
		return err
	}
	domainDir := baseDirectory + "/" + directoryName
	for _, name := range fileNames {
		if err := utl.CreateFileWithContent(domainDir, name, generalFileContent); err != nil {
			return err
		}
	}
	for _, name := range testFileNames {
		if err := utl.CreateFileWithContent(domainDir, name, testFileContent); err != nil {
			return err
		}
	}
	return nil
}
