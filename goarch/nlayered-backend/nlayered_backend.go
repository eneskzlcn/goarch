package nlayered_backend

import (
	"github.com/eneskzlcn/goarch/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch/common"
)

type NLayeredBackend struct {
	arch *arch.Architecture
}

func New(basePath string) *NLayeredBackend {
	arch := arch.New(basePath)
	nLayeredBackend := &NLayeredBackend{arch: arch}
	nLayeredBackend.initializeArchitecture()
	return nLayeredBackend
}
func (nlb *NLayeredBackend) Create() error {
	return nlb.arch.Create()
}
func (nlb *NLayeredBackend) addInitializedCommonRootDirectories() {
	nlb.arch.PutDirectory(".dev", common.DevDirectory)
	nlb.arch.PutDirectory(".cd", common.CdDirectory)
	nlb.arch.PutDirectory("seed", common.SeedDirectory)
	nlb.arch.PutDirectory("cmd", common.CmdDirectory)
}
func (nlb *NLayeredBackend) initializeArchitecture() {
	nlb.addInitializedCommonRootDirectories()
	nlb.addInitializedNLayeredBackendSpecificDirectories()
}
func (nlb *NLayeredBackend) addInitializedNLayeredBackendSpecificDirectories() {
	nlb.arch.PutDirectory("internal", internalDirectory)
}
