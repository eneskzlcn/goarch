package file

import (
	"fmt"
	"github.com/eneskzlcn/goarch/goarch/fileutil"
)

type MarkdownFile struct {
	Content string
}

func NewMarkdownFile(content string) MarkdownFile {
	return MarkdownFile{Content: content}
}
func (m MarkdownFile) Create(path, name string) error {
	filename := m.nameToFilename(name)
	if m.Content == "" {
		return m.createEmptyFile(path, filename)
	}
	return fileutil.CreateFileWithContent(path, filename, m.Content)
}

func (m MarkdownFile) createEmptyFile(directoryPath, fileName string) error {
	content := m.emptyFileContent(directoryPath)
	return fileutil.CreateFileWithContent(directoryPath, fileName, content)
}

func (m MarkdownFile) nameToFilename(name string) string {
	if fileutil.IsMarkdownFilename(name) {
		return name
	}
	return name + ".md"
}

func (m MarkdownFile) emptyFileContent(directoryPath string) string {
	return fmt.Sprintf("[//]: # %s", directoryPath)
}
