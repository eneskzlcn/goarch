package microservice

import (
	"github.com/eneskzlcn/goarch/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch/common"
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
	m.addInitializedCommonDirectories()
	m.addInitializedMicroserviceSpecificDirectories()
}
func (m *Microservice) addInitializedCommonDirectories() {
	m.arch.PutDirectory("seed", common.SeedDirectory)
	m.arch.PutDirectory(".dev", common.DevDirectory)
	m.arch.PutDirectory(".cd", common.CdDirectory)
	m.arch.PutDirectory("cmd", common.CmdDirectory)
	m.arch.PutDirectory("logger", common.LoggerDirectory)
	m.arch.PutDirectory("postgres", common.PostgresDirectory)
	m.arch.PutDirectory("rabbitmq", common.RabbitmqDirectory)

}
func (m *Microservice) addInitializedMicroserviceSpecificDirectories() {
	m.arch.PutDirectory("internal", internalDirectory)
}
