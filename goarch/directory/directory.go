package directory

import "github.com/eneskzlcn/goarch/architecture"

type CreatorByArchitecture = func(p architecture.Type) architecture.Directory

func FindPathOfGivenDirectoryByNameAndArchitecture(directoryName string, architectureType architecture.Type) string {
	directoryNameToPathMapOfGivenArchitecture := architectureToDirectoryNamesThenDirectoryNameToPath[architectureType]
	pathOfGivenDirectoryName := directoryNameToPathMapOfGivenArchitecture[directoryName]
	return pathOfGivenDirectoryName
}

type NameToPathMap map[string]string

func CreateDirectoriesWithEmptyGoAndTestFiles(dirPath string, dirNames ...string) architecture.Directories {
	directories := make(architecture.Directories, 0)
	for _, name := range dirNames {
		directories = append(directories, architecture.Directory{
			AbsPath: dirPath + "/",
			Name:    name,
			SubDir:  nil,
			Files: architecture.Files{
				{
					Name:    name + ".go",
					Content: "package " + name,
				},
				{
					Name:    name + "_test.go",
					Content: "package " + name + "_test",
				},
			},
		})
	}
	return directories
}

var architectureToDirectoryNamesThenDirectoryNameToPath = map[architecture.Type]NameToPathMap{
	architecture.Microservice: {
		"kafka":    architecture.RootDirectory,
		"logger":   architecture.RootDirectory,
		"postgres": architecture.RootDirectory,
		"rabbitmq": architecture.RootDirectory,
		"server":   architecture.RootDirectory,
		"config":   architecture.InternalDirectory,
		"client":   architecture.InternalDirectory,
		"mocks":    architecture.InternalDirectory,
		"util":     architecture.InternalDirectory,
	},
	architecture.NLayeredWebApp: {
		"kafka":    architecture.CoreDirectory,
		"logger":   architecture.CoreDirectory,
		"postgres": architecture.CoreDirectory,
		"rabbitmq": architecture.CoreDirectory,
		"server":   architecture.CoreDirectory,
		"config":   architecture.InternalDirectory,
		"client":   architecture.CoreDirectory,
		"mocks":    architecture.InternalDirectory,
		"util":     architecture.CoreDirectory,
	},
	architecture.NLayeredBackend: {
		"kafka":    architecture.CoreDirectory,
		"logger":   architecture.CoreDirectory,
		"postgres": architecture.CoreDirectory,
		"rabbitmq": architecture.CoreDirectory,
		"server":   architecture.CoreDirectory,
		"config":   architecture.InternalDirectory,
		"client":   architecture.CoreDirectory,
		"mocks":    architecture.InternalDirectory,
		"util":     architecture.CoreDirectory,
	},
}
