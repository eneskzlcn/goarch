package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/architecture"
)

//MARK: All common root directories including .dev, seed, cmd inside of that variable RootDirectories

var RootDirectories = architecture.Directories{
	seedDirectory,
	devDirectory,
	cmdDirectory,
	cdDirectory(),
}

//MARK: ./seed directory initialization

//go:embed templates/seed/cmd/main_go.arch
var seedCmdMainFileContent string

//go:embed templates/seed/readme_md.arch
var seedReadmeMDFileContent string

//go:embed templates/seed/seed_go.arch
var seedGoFileContent string

//go:embed templates/seed/seed_test.arch
var seedTestFileContent string

//go:embed templates/seed/create_sql.arch
var seedCreateSeedSqlFileContent string

//go:embed templates/seed/drop_sql.arch
var seedDropSeedSqlFileContent string

var seedDirectory = architecture.Directory{
	AbsPath: architecture.RootDirectory,
	Name:    "seed",
	SubDir: architecture.Directories{
		{
			AbsPath: "./seed",
			Name:    "cmd",
			SubDir:  nil,
			Files: architecture.Files{
				{
					Name:    "main.go",
					Content: seedCmdMainFileContent,
				},
			},
		},
	},
	Files: architecture.Files{
		{
			Name:    "create-seed.sql",
			Content: seedCreateSeedSqlFileContent,
		},
		{
			Name:    "drop-seed.sql",
			Content: seedDropSeedSqlFileContent,
		},
		{
			Name:    "seed.go",
			Content: seedGoFileContent,
		},
		{
			Name:    "seed_test.go",
			Content: seedTestFileContent,
		},
	},
}

//MARK: ./.dev directory initialization.

//go:embed templates/.dev/default_config_yaml.arch
var devDefaultConfigYamlContent string

var devDirectory = architecture.Directory{
	AbsPath: architecture.RootDirectory,
	Name:    ".dev",
	SubDir:  nil,
	Files: architecture.Files{
		{
			Name:    "dev.yaml",
			Content: devDefaultConfigYamlContent,
		},
		{
			Name:    "qa.yaml",
			Content: devDefaultConfigYamlContent,
		},
		{
			Name:    "prod.yaml",
			Content: devDefaultConfigYamlContent,
		},
		{
			Name:    "test.yaml",
			Content: devDefaultConfigYamlContent,
		},
	},
}

// MARK: ./cmd directory initialization.

//go:embed templates/cmd/main_go.arch
var cmdDirectoryMainFileContent string

var cmdDirectory = architecture.Directory{
	AbsPath: architecture.RootDirectory,
	Name:    "cmd",
	SubDir:  nil,
	Files: architecture.Files{
		{
			Name:    "main.go",
			Content: cmdDirectoryMainFileContent,
		},
	},
}

//MARK: ./.cd directory initialization

//go:embed templates/continous-delivery/ci-cd_yaml.arch
var cdYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/service_yaml.arch
var cdServiceYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/deployment_yml.arch
var cdDeploymentYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/pv_yaml.arch
var cdPvYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/pv-claim_arch
var cdPvClaimFileContent string

func cdDirectory() architecture.Directory {
	name := ".cd"
	cdDir := architecture.RootDirectory + "/" + name
	return architecture.Directory{
		AbsPath: architecture.RootDirectory,
		Name:    name,
		SubDir: architecture.Directories{
			{
				AbsPath: cdDir,
				Name:    "deployment-artifacts",
				SubDir:  nil,
				Files: architecture.Files{
					{
						Name:    "service.yaml",
						Content: cdServiceYamlFileContent,
					},
					{
						Name:    "pv.yaml",
						Content: cdPvYamlFileContent,
					},
					{
						Name:    "pv-claim.yaml",
						Content: cdPvClaimFileContent,
					},
				},
			},
		},
		Files: architecture.Files{
			{
				Name:    "cd.yaml",
				Content: cdYamlFileContent,
			},
		},
	}
}
