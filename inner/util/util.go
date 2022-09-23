package util

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed ctxutil_go.arch
var ctxUtilFileContent string

//go:embed ctxutil_test_go.arch
var ctxUtilTestFileContent string

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   "./internal",
	arch.NLayeredWebApp: "./internal/core",
}

const (
	directoryName       = "util"
	ctxUtilDirName      = "ctxutil"
	ctxUtilFileName     = "ctxutil.go"
	ctxUtilTestFileName = "ctxutil_test.go"
)

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}
	utilDir := baseDir + "/" + directoryName
	if err := utl.CreateDirectory(utilDir, ctxUtilDirName); err != nil {
		return err
	}
	ctxutilDir := utilDir + "/" + ctxUtilDirName

	if err := utl.CreateFileWithContent(ctxutilDir, ctxUtilFileName, ctxUtilFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(ctxutilDir, ctxUtilTestFileName, ctxUtilTestFileContent); err != nil {
		return err
	}
	return nil
}
