package file

import (
	"github.com/eneskzlcn/goarch/goarch/fileutil"
	"path"
)

type GoTestFile struct {
	Content string
}

func (g GoTestFile) Create(directoryPath, name string) error {
	fileName := g.nameToFilename(name)
	if g.Content == "" {
		return g.createEmptyFile(directoryPath, fileName)
	}
	return fileutil.CreateFileWithContent(directoryPath, fileName, g.Content)
}
func (g GoTestFile) createEmptyFile(directoryPath, fileName string) error {
	packageName := path.Base(directoryPath)
	fileContent := g.emptyFileContent(packageName)
	return fileutil.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (g GoTestFile) nameToFilename(name string) string {
	return name + "_test.go"
}
func (g GoTestFile) emptyFileContent(packageName string) string {
	return "package " + packageName + "_test"
}
