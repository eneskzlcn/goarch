package nlayered_web

import (
	"github.com/eneskzlcn/goarch/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch/common"
)

type NLayeredWeb struct {
	arch *arch.Architecture
}

func New(basePath string) *NLayeredWeb {
	arch := arch.New(basePath)
	nlw := NLayeredWeb{arch: arch}
	nlw.initializeArchitecture()
	return &nlw
}
func (nlw *NLayeredWeb) initializeArchitecture() {
	nlw.addInitializedCommonRootDirectories()
	nlw.addInitializedNLayeredWebSpecificDirectories()
}
func (nlw *NLayeredWeb) addInitializedCommonRootDirectories() {
	nlw.arch.PutDirectory(".dev", common.DevDirectory)
	nlw.arch.PutDirectory(".cd", common.CdDirectory)
	nlw.arch.PutDirectory("seed", common.SeedDirectory)
	nlw.arch.PutDirectory("cmd", common.CmdDirectory)
}
func (nlw *NLayeredWeb) addInitializedNLayeredWebSpecificDirectories() {
	nlw.arch.PutDirectory("internal", internalDirectory)
}
