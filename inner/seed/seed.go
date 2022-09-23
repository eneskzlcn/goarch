package seed

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed create_sql.arch
var createSqlFileContent string

//go:embed drop_sql.arch
var dropSqlFileContent string

//go:embed seed_go.arch
var seedFileContent string

//go:embed seed_test.arch
var seedTestFileContent string

//go:embed main_go.arch
var mainFileContent string

//go:embed readme_md.arch
var readmeFileContent string

const (
	directoryName        = "seed"
	commandDirectoryName = "cmd"

	mainFileName      = "main.go"
	seedFileName      = "seed.go"
	seedTestFileName  = "seed_test.go"
	createSqlFileName = "create_seed.sql"
	dropSqlFileName   = "drop_seed.sql"
	readmeFileName    = "README.md"
)

var architectureAbsPathMap = map[arch.Type]string{
	arch.Microservice:   ".",
	arch.NLayeredWebApp: "./internal/core",
}

func PrepareDirectory(architecture arch.Type) error {
	baseDir := architectureAbsPathMap[architecture]
	if err := utl.CreateDirectory(baseDir, directoryName); err != nil {
		return err
	}

	seedDir := baseDir + "/" + directoryName
	if err := utl.CreateFileWithContent(seedDir, seedFileName, seedFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(seedDir, seedTestFileName, seedTestFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(seedDir, createSqlFileName, createSqlFileContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(seedDir, dropSqlFileName, dropSqlFileContent); err != nil {
		return err
	}

	if err := utl.CreateDirectory(seedDir, commandDirectoryName); err != nil {
		return err
	}
	cmdDir := seedDir + "/" + commandDirectoryName
	if err := utl.CreateFileWithContent(cmdDir, mainFileName, mainFileContent); err != nil {
		return err
	}
	return nil
}
