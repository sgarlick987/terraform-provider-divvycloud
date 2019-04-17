package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sgarlick987/terraform-provider-divvycloud/divvycloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: divvycloud.Provider})
}
