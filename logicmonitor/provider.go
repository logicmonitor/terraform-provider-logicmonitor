package logicmonitor

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//Provider for LogicMonitor
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_ID", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_KEY", nil),
			},
			"company": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_COMPANY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"logicmonitor_collector_group": resourceCollectorGroup(),
			"logicmonitor_device":          resourceDevices(),
			"logicmonitor_device_group":    resourceDeviceGroup(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"logicmonitor_device_group": dataSourceFindDeviceGroups(),
			"logicmonitor_collectors":   dataSourceFindCollectors(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessID:  d.Get("api_id").(string),
		AccessKey: d.Get("api_key").(string),
		Company:   d.Get("company").(string),
	}
	log.Println("[INFO] Initializing LM client")
	return config.NewLMClient()
}
