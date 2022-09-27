package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: initialize mocks directory.

//go:embed templates/mocks/domain/mock_domain_repository.go
var mockDomainRepositoryFileContent string

var MocksDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"domain": directory.Directory{
			Files: file.Files{
				"mock_domain_repository": file.NewGoFile(mockDomainRepositoryFileContent),
			},
		},
	},
}
