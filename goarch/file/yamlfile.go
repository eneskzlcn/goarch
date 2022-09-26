package file

import "github.com/eneskzlcn/goarch/goarch/fileutil"

type YamlFile struct {
	Content string
}

func NewYamlFile(content string) YamlFile {
	return YamlFile{Content: content}
}
func (y YamlFile) Create(path, name string) error {
	fileName := y.nameToFilename(name)
	if y.Content == "" {
		return y.createEmptyFile(path, fileName)
	}
	return fileutil.CreateFileWithContent(path, fileName, y.Content)
}
func (y YamlFile) createEmptyFile(directoryPath, fileName string) error {
	fileContent := y.emptyFileContent(fileName)
	return fileutil.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (y YamlFile) emptyFileContent(fileName string) string {
	return "--- #" + fileName
}
func (y YamlFile) nameToFilename(name string) string {
	return name + ".yaml"
}
