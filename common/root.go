package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/arch"
)

//MARK: All common root directories including .dev, seed, cmd inside of that variable RootDirectories

var RootDirectories = arch.Directories{
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

var seedDirectory = arch.Directory{
	AbsPath: arch.RootDirectory,
	Name:    "seed",
	SubDir: arch.Directories{
		{
			AbsPath: "./seed",
			Name:    "cmd",
			SubDir:  nil,
			Files: arch.Files{
				{
					Name:    "main.go",
					Content: seedCmdMainFileContent,
				},
			},
		},
	},
	Files: arch.Files{
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

var devDirectory = arch.Directory{
	AbsPath: arch.RootDirectory,
	Name:    ".dev",
	SubDir:  nil,
	Files: arch.Files{
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

var cmdDirectory = arch.Directory{
	AbsPath: arch.RootDirectory,
	Name:    "cmd",
	SubDir:  nil,
	Files: arch.Files{
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

func cdDirectory() arch.Directory {
	name := ".cd"
	cdDir := arch.RootDirectory + "/" + name
	return arch.Directory{
		AbsPath: arch.RootDirectory,
		Name:    name,
		SubDir: arch.Directories{
			{
				AbsPath: cdDir,
				Name:    "deployment-artifacts",
				SubDir:  nil,
				Files: arch.Files{
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
		Files: arch.Files{
			{
				Name:    "cd.yaml",
				Content: cdYamlFileContent,
			},
		},
	}
}
