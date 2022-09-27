package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: postgres directory initialization.

//go:embed templates/postgres/postgres_go.arch
var postgresGoFileContent string

//go:embed templates/postgres/postgres_test.arch
var postgresGoTestFileContent string

//go:embed templates/postgres/mock_postgres_go.arch
var mockPostgresGoFileContent string

var PostgresDirectory = directory.Directory{
	Files: file.Files{
		"postgres":      file.NewGoFile(postgresGoFileContent),
		"postgres_test": file.NewGoTestFile(postgresGoTestFileContent),
		"mock_postgres": file.NewGoFile(mockPostgresGoFileContent),
	},
}
