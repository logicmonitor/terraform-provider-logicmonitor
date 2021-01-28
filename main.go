package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/logicmonitor/terraform-provider-logicmonitor/logicmonitor"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logicmonitor.Provider,
	})
}
