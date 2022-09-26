package arch

import "github.com/eneskzlcn/goarch/goarch/directory"

type Architecture struct {
	BasePath    string
	Directories directory.Directories
}

func New(basePath string) *Architecture {
	return &Architecture{BasePath: basePath, Directories: directory.Directories{}}
}
func (a *Architecture) Create() error {
	for name, directory := range a.Directories {
		if err := directory.Create(a.BasePath, name); err != nil {
			return err
		}
	}
	return nil
}
func (a *Architecture) PutDirectory(name string, directory directory.Directory) {
	a.Directories[name] = directory
}

type Microservice struct {
	Architecture
}
