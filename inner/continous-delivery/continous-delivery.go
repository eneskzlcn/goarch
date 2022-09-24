package continous_delivery

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/inner/core/tech"
	"github.com/eneskzlcn/goarch/inner/core/utl"
)

//go:embed ci-cd_yaml.arch
var cdYamlContent string

//go:embed deployment_yml.arch
var deploymentYmlContent string

//go:embed pv-claim_arch
var pvClaimYmlContent string

//go:embed pv_yaml.arch
var pvYmlContent string

//go:embed service_yaml.arch
var serviceYmlContent string

const (
	defaultDirectory                 = "."
	directoryName                    = ".cd"
	cdYamlDirectory                  = defaultDirectory + "/" + directoryName
	cdYamlFileName                   = "cd.yaml"
	deploymentArtifactsDirectoryName = "deployment-artifacts"
	deploymentArtifactsDirectory     = cdYamlDirectory + "/" + deploymentArtifactsDirectoryName
	deploymentYmlName                = "deployment.yml"
	serviceYmlName                   = "service.yml"
	pvYmlName                        = "pv.yml"
	pvClaimYmlName                   = "pv-claim.yml"
)

func PrepareDirectory(options tech.Options) error {
	if err := utl.CreateDirectory(defaultDirectory, directoryName); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(cdYamlDirectory, cdYamlFileName, cdYamlContent); err != nil {
		return err
	}

	if err := utl.CreateDirectory(cdYamlDirectory, deploymentArtifactsDirectoryName); err != nil {
		return err
	}

	if err := utl.CreateFileWithContent(deploymentArtifactsDirectory,
		deploymentYmlName, deploymentYmlContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(deploymentArtifactsDirectory,
		serviceYmlName, serviceYmlContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(deploymentArtifactsDirectory, pvYmlName, pvYmlContent); err != nil {
		return err
	}
	if err := utl.CreateFileWithContent(deploymentArtifactsDirectory, pvClaimYmlName, pvClaimYmlContent); err != nil {
		return err
	}
	return nil
}
