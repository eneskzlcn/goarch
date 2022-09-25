package microservice

import (
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/common"
)

var InternalDirectory = architecture.Directory{
	AbsPath: architecture.RootDirectory,
	Name:    "internal",
	SubDir: architecture.Directories{
		common.ConfigDirectory(architecture.Microservice),
		domainDirectory,
		common.ClientDirectory(architecture.Microservice),
		common.UtilDirectory(architecture.Microservice),
		common.MocksDirectory(architecture.Microservice),
	},
	Files: nil,
}

func domainFiles() architecture.Files {
	packageName := "domain"
	files := make(architecture.Files, 0)
	domainFileNames := []string{"request", "response", "domain", "service", "handler", "repository"}
	for _, name := range domainFileNames {
		file := architecture.File{
			Name:    name + ".go",
			Content: "package " + packageName,
		}
		testFile := architecture.File{
			Name:    name + "_test" + ".go",
			Content: "package " + packageName + "_test",
		}
		files = append(files, file, testFile)
	}
	return files
}

var domainDirectory = architecture.Directory{
	AbsPath: architecture.InternalDirectory,
	Name:    "domain",
	SubDir:  nil,
	Files:   domainFiles(),
}
