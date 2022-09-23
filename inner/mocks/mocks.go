package mocks

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

const (
	directoryName              = "mocks"
	baseDir                    = "./internal"
	userMockDirName            = "user"
	mockUserRepositoryFileName = "mock_user_repository.go"
)

//go:embed mock_user_repository_go.arch
var mockUserRepositoryContent string

func PrepareDirectory() error {
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}
	mocksDir := baseDir + "/" + directoryName
	if err := utl.CreateDirectory(mocksDir, userMockDirName); err != nil {
		return err
	}
	userDir := mocksDir + "/" + userMockDirName
	if err := utl.CreateFileWithContent(userDir, mockUserRepositoryFileName, mockUserRepositoryContent); err != nil {
		return err
	}
	return nil
}
