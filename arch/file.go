package arch

import (
	"github.com/eneskzlcn/goarch/utl"
)

type Files []File

type File struct {
	Name    string
	Content string
}

func (f Files) Create(to string) error {
	for _, file := range f {
		if err := file.Create(to); err != nil {
			return err
		}
	}
	return nil
}

func (f File) Create(to string) error {
	return utl.CreateFileWithContent(to, f.Name, f.Content)
}
