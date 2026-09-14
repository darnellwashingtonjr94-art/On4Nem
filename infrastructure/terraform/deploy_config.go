package terraform

import "log"

type TerraformDeployConfig struct {
	Region   string
	Instance string
}

func GenerateDeploymentScript() {
	log.Println("Generating Terraform IaC configurations for container container clusters...")
}
