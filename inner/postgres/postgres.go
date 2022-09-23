package postgres

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/arch"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed postgres_go.arch
var postgresFileContent string

//go:embed postgres_test.arch
var postgresTestFileContent string

//go:embed mock_postgres_go.arch
var mockPostgresFileContent string

const (
	postgresFileName     = "postgres.go"
	directoryName        = "postgres"
	postgresTestFileName = "postgres_test.go"
	mockPostgresFileName = "mock_postgres.go"
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

	postgresDir := baseDir + "/" + directoryName
	if err := utl.CreateFileWithContent(postgresDir, postgresFileName, postgresFileContent); err != nil {
		return err
	}

	if err := utl.CreateFileWithContent(postgresDir, postgresTestFileName, postgresTestFileContent); err != nil {
		return err
	}

	if err := utl.CreateFileWithContent(postgresDir, mockPostgresFileName, mockPostgresFileContent); err != nil {
		return err
	}

	return nil
}
