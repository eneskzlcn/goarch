package file

import (
	"fmt"
	"github.com/eneskzlcn/goarch/goarch/fileutil"
	"path"
)

type GoHtmlFile struct {
	Content string
}

func NewGoHtmlFile(content string) GoHtmlFile {
	return GoHtmlFile{Content: content}
}

func (g GoHtmlFile) Create(directoryPath, name string) error {
	fileName := g.nameToFilename(name)
	if g.Content == "" {
		return g.createEmptyFile(directoryPath, fileName)
	}
	return fileutil.CreateFileWithContent(directoryPath, fileName, g.Content)
}
func (g GoHtmlFile) createEmptyFile(directoryPath, fileName string) error {
	packageName := path.Base(directoryPath)
	fileContent := g.emptyFileContent(packageName)
	return fileutil.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (g GoHtmlFile) nameToFilename(name string) string {
	if fileutil.IsGohtmlFilename(name) {
		return name
	}
	return name + ".gohtml"
}
func (g GoHtmlFile) emptyFileContent(packageName string) string {
	return fmt.Sprintf("<div>%s</div>", packageName)
}
