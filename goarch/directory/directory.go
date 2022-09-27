package directory

import (
	"github.com/eneskzlcn/goarch/goarch/file"
	"github.com/eneskzlcn/goarch/goarch/fileutil"
	"github.com/eneskzlcn/goarch/goarch/pathutil"
)

type Directory struct {
	SubDirs Directories
	Files   file.Files
}

func New() Directory {
	return Directory{SubDirs: Directories{}, Files: file.Files{}}
}

type Directories map[string]Directory

type NameFilesMap map[string]file.Files

func (d Directory) Create(path, directoryName string) error {
	if err := fileutil.CreateDirectory(path, directoryName); err != nil {
		return err
	}
	childDirectoriesAndFilesPath := pathutil.ExtendPath(path, directoryName)
	if err := d.createSubDirectories(childDirectoriesAndFilesPath); err != nil {
		return err
	}
	if err := d.createFiles(childDirectoriesAndFilesPath); err != nil {
		return err
	}
	return nil
}
func (d Directory) createFiles(filesPath string) error {
	for name, file := range d.Files {
		if err := file.Create(filesPath, name); err != nil {
			return err
		}
	}
	return nil
}
func (d Directory) PutSubDirectory(name string, directory Directory) {
	d.SubDirs[name] = directory
}
func (d Directory) PutFile(name string, file file.File) {
	d.Files[name] = file
}
func (d Directory) createSubDirectories(subDirectoryPath string) error {
	for name, directory := range d.SubDirs {
		if err := directory.Create(subDirectoryPath, name); err != nil {
			return err
		}
	}
	return nil
}
