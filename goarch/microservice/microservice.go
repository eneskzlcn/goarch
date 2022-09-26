package microservice

import (
	"github.com/eneskzlcn/goarch/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch/directory"
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
	m.arch.PutDirectory(".dev", m.initializeDirectory(".dev"))
}
func (m *Microservice) initializeLoggerDirectory() {
	panic("implement me")
}
func (m *Microservice) initializeDirectory(directoryName string) directory.Directory {
	devDirectory := directory.New()
	nameFileMapOfDirectory := directoryToNameFileMapper[directoryName]
	for name, file := range nameFileMapOfDirectory {
		devDirectory.PutFile(name, file)
	}
	return devDirectory
}
