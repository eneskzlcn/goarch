package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: .cd directory initialization.

//go:embed templates/continous-delivery/ci-cd_yaml.arch
var cdYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/service_yaml.arch
var serviceYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/deployment_yml.arch
var deploymentYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/pv_yaml.arch
var pvYamlFileContent string

//go:embed templates/continous-delivery/deployment_artifacts/pv-claim_arch
var pvClaimYamlFileContent string

var CdDirectory = directory.Directory{
	SubDirs: directory.Directories{
		"deployment-artifacts": {
			Files: file.Files{
				"service":    file.NewYamlFile(serviceYamlFileContent),
				"deployment": file.NewYamlFile(deploymentYamlFileContent),
				"pv":         file.NewYamlFile(pvYamlFileContent),
				"pv-claim":   file.NewYamlFile(pvClaimYamlFileContent),
			},
		},
	},
	Files: file.Files{
		"cd": file.NewYamlFile(cdYamlFileContent),
	},
}
