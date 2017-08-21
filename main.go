package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-logicmonitor/logicmonitor"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logicmonitor.Provider,
	})
}
