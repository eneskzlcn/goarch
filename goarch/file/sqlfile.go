package file

import (
	"fmt"
	"github.com/eneskzlcn/goarch/goarch/fileutil"
)

type SqlFile struct {
	Content string
}

func NewSqlFile(content string) SqlFile {
	return SqlFile{Content: content}
}
func (s SqlFile) Create(path, name string) error {
	filename := s.nameToFilename(name)
	if s.Content == "" {
		return s.createEmptyFile(path, filename)
	}
	return fileutil.CreateFileWithContent(path, filename, s.Content)
}

func (s SqlFile) createEmptyFile(directoryPath, fileName string) error {
	fileContent := s.emptyFileContent(directoryPath)
	return fileutil.CreateFileWithContent(directoryPath, fileName, fileContent)
}

func (s SqlFile) nameToFilename(name string) string {
	if fileutil.IsSqlFilename(name) {
		return name
	}
	return name + ".sql"
}

func (s SqlFile) emptyFileContent(directoryPath string) string {
	return fmt.Sprintf("-- %s\n%s", directoryPath, "-- SELECT * from users")
}
