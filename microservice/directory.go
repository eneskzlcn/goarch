package microservice

import (
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/common"
)

var InternalDirectory = arch.Directory{
	AbsPath: arch.RootDirectory,
	Name:    "internal",
	SubDir: arch.Directories{
		common.ConfigDirectory(architecture),
		domainDirectory,
		common.ClientDirectory(architecture),
		common.UtilDirectory(architecture),
		common.MocksDirectory(architecture),
	},
	Files: nil,
}

func domainFiles() arch.Files {
	packageName := "domain"
	files := make(arch.Files, 0)
	domainFileNames := []string{"request", "response", "domain", "service", "handler", "repository"}
	for _, name := range domainFileNames {
		file := arch.File{
			Name:    name + ".go",
			Content: "package " + packageName,
		}
		testFile := arch.File{
			Name:    name + "_test" + ".go",
			Content: "package " + packageName + "_test",
		}
		files = append(files, file, testFile)
	}
	return files
}

var domainDirectory = arch.Directory{
	AbsPath: arch.InternalDirectory,
	Name:    "domain",
	SubDir:  nil,
	Files:   domainFiles(),
}
