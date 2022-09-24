package anotherway

import (
	"fmt"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

type Directories []Directory

type Directory struct {
	AbsPath string
	Name    string
	SubDir  Directories
	Files   Files
}

func (d *Directories) Create() error {
	for _, dir := range *d {
		if err := dir.Create(); err != nil {
			return err
		}
	}
	return nil
}
func (d *Directory) Create() error {
	if err := utl.CreateDirectory(d.AbsPath, d.Name); err != nil {
		return err
	}
	for _, file := range d.Files {
		if err := file.Create(d.path()); err != nil {
			return err
		}
	}
	return d.SubDir.Create()
}
func (d *Directory) path() string {
	return fmt.Sprintf("%s/%s", d.AbsPath, d.Name)
}
