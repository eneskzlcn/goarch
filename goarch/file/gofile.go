package file

import (
	"github.com/eneskzlcn/goarch/goarch/fileutil"
	"path"
)

type GoFile struct {
	Content string
}

func NewGoFile(content string) GoFile {
	return GoFile{Content: content}
}

func (g GoFile) Create(directoryPath, name string) error {
	fileName := g.nameToFilename(name)
	if g.Content == "" {
		return g.createEmptyFile(directoryPath, fileName)
	}
	return fileutil.CreateFileWithContent(directoryPath, fileName, g.Content)
}
func (g GoFile) createEmptyFile(directoryPath, fileName string) error {
	packageName := path.Base(directoryPath)
	fileContent := g.emptyFileContent(packageName)
	return fileutil.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (g GoFile) nameToFilename(name string) string {
	return name + ".go"
}
func (g GoFile) emptyFileContent(packageName string) string {
	return "package " + packageName
}
