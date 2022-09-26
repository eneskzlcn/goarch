package file

type File interface {
	Create(path, name string) error
	createEmptyFile(directoryPath, fileName string) error
	nameToFilename(name string) string
	emptyFileContent(directoryPath string) string
}

type Files map[string]File
