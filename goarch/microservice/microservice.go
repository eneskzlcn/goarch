package microservice

import (
	"github.com/eneskzlcn/goarch/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

type Microservice struct {
	arch *arch.Architecture
}

func New(basePath string) *Microservice {
	arch := arch.New(basePath)
	ms := Microservice{arch: arch}
	ms.initializeArchitecture()
	return &ms
}
func (m *Microservice) Create() error {
	return m.arch.Create()
}
func (m *Microservice) initializeArchitecture() {
	m.initializeDevDirectory()
}
func (m *Microservice) initializeDevDirectory() {
	devDirectory := directory.New()
	devDirectoryFileNames := []string{"local", "dev", "qa", "prod"}
	devDirectoryFileTemplate := file.YamlFile{}
	for _, name := range devDirectoryFileNames {
		devDirectory.PutFile(name, devDirectoryFileTemplate)
	}
	m.arch.PutDirectory(".dev", devDirectory)
}
