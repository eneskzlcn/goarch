package arch

import (
	"fmt"
	"github.com/eneskzlcn/goarch/utl"
)

type Directories []Directory

type Directory struct {
	AbsPath string
	Name    string
	SubDir  Directories
	Files   Files
}

func (d Directories) Create() error {
	for _, dir := range d {
		if err := dir.Create(); err != nil {
			return err
		}
	}
	return nil
}
func (p Directories) Print() {
	fmt.Println()
	for _, dir := range p {
		dir.Print()
	}
}
func (d Directory) Create() error {
	fmt.Printf("Directory created to %s with name %s \n", d.AbsPath, d.Name)
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
func (d Directory) path() string {
	return fmt.Sprintf("%s/%s", d.AbsPath, d.Name)
}
func (d Directory) Print() {
	fmt.Println(d.Name)
	for _, dir := range d.SubDir {
		fmt.Printf("\t %s/%s", dir.path(), dir.Name)
	}
	fmt.Println()
}
