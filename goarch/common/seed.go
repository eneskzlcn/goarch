package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: seed directory initialization.

//go:embed templates/seed/seed_go.arch
var seedSeedGoFileContent string

//go:embed templates/seed/seed_test.arch
var seedSeedTestGoFileContent string

//go:embed templates/seed/create_sql.arch
var seedCreateSeedSqlFileContent string

//go:embed templates/seed/drop_sql.arch
var seedDropSeedSqlFileContent string

//go:embed templates/seed/readme_md.arch
var seedReadmeMarkdownFileContent string

//go:embed templates/seed/cmd/main_go.arch
var seedCmdMainGoFileContent string

var SeedDirectory = directory.Directory{
	SubDirs: map[string]directory.Directory{
		"cmd": {
			Files: file.Files{
				"main": file.NewGoFile(seedCmdMainGoFileContent),
			},
		},
	},
	Files: file.Files{
		"create_seed": file.NewSqlFile(seedCreateSeedSqlFileContent),
		"drop_seed":   file.NewSqlFile(seedDropSeedSqlFileContent),
		"readme":      file.NewMarkdownFile(seedReadmeMarkdownFileContent),
		"seed":        file.NewGoFile(seedSeedGoFileContent),
		"seed_test":   file.NewGoTestFile(seedSeedTestGoFileContent),
	},
}
