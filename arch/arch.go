package arch

import (
	"github.com/eneskzlcn/goarch/architecture"
	"github.com/eneskzlcn/goarch/utl"
	"path"
)

type File interface {
	Create(path, name string) error
}
type GoFile struct {
	Content string
}

func (g GoFile) Create(directoryPath, name string) error {
	fileName := g.nameToFilename(name)
	if g.Content == "" {
		return g.createEmptyFile(directoryPath, fileName)
	}
	return utl.CreateFileWithContent(directoryPath, fileName, g.Content)
}
func (g GoFile) createEmptyFile(directoryPath, fileName string) error {
	packageName := path.Base(directoryPath)
	fileContent := g.emptyFileContent(packageName)
	return utl.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (g GoFile) nameToFilename(name string) string {
	return name + ".go"
}
func (g GoFile) emptyFileContent(packageName string) string {
	return "package " + packageName
}

type YamlFile struct {
	Content string
}

func (y YamlFile) Create(path, name string) error {
	fileName := y.nameToFileName(name)
	if y.Content == "" {
		return y.createEmptyFile(path, fileName)
	}
	return utl.CreateFileWithContent(path, fileName, y.Content)
}
func (y YamlFile) createEmptyFile(directoryPath, fileName string) error {
	fileContent := y.emptyFileContent(fileName)
	return utl.CreateFileWithContent(directoryPath, fileName, fileContent)
}
func (y YamlFile) emptyFileContent(fileName string) string {
	return "--- #" + fileName
}
func (y YamlFile) nameToFileName(name string) string {
	return name + ".yaml"
}

type GoTestFile struct {
}

type Directory struct {
	Directories NameDirectoryMap
	Files       NameFileMap
}

func (d Directory) Create(path, directoryName string) error {
	if err := utl.CreateDirectory(path, directoryName); err != nil {
		return err
	}
	childDirectoriesAndFilesPath := ExtendPath(path, directoryName)
	if err := d.CreateSubDirectories(childDirectoriesAndFilesPath); err != nil {
		return err
	}
	if err := d.CreateFiles(childDirectoriesAndFilesPath); err != nil {
		return err
	}
	return nil
}
func (d Directory) CreateFiles(filesPath string) error {
	for name, directory := range d.Directories {
		if err := directory.Create(filesPath, name); err != nil {
			return err
		}
	}
	return nil
}
func (d Directory) CreateSubDirectories(subDirectoryPath string) error {
	for name, directory := range d.Directories {
		if err := directory.Create(subDirectoryPath, name); err != nil {
			return err
		}
	}
	return nil
}
func ExtendPath(path, name string) string {
	return path + "/" + name
}

type Architecture struct {
	Type        architecture.Type
	BasePath    string
	Directories NameDirectoryMap
	Files       NameFileMap
}

func (a *Architecture) Create() error {
	for name, directory := range a.Directories {
		if err := directory.Create(a.BasePath, name); err != nil {
			return err
		}
	}
	return nil
}

type NameDirectoryMap map[string]Directory

type NameFileMap map[string]File

type Microservice struct {
	Architecture
}

func Try() {
	m := Microservice{}
	m.Directories[".dev"] = Directory{
		Files: nil,
	}
	m.Create()
}
