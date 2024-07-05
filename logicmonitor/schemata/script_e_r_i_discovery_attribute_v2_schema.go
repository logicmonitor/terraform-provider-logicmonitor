package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ScriptERIDiscoveryAttributeV2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"groovy_script": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"linux_cmdline": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"linux_script": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"win_cmdline": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"win_script": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetScriptERIDiscoveryAttributeV2SubResourceData(m []*models.ScriptERIDiscoveryAttributeV2) (d []*map[string]interface{}) {
	for _, scriptERIDiscoveryAttributeV2 := range m {
		if scriptERIDiscoveryAttributeV2 != nil {
			properties := make(map[string]interface{})
			properties["groovy_script"] = scriptERIDiscoveryAttributeV2.GroovyScript
			properties["linux_cmdline"] = scriptERIDiscoveryAttributeV2.LinuxCmdline
			properties["linux_script"] = scriptERIDiscoveryAttributeV2.LinuxScript
			properties["name"] = scriptERIDiscoveryAttributeV2.Name
			properties["type"] = scriptERIDiscoveryAttributeV2.Type
			properties["win_cmdline"] = scriptERIDiscoveryAttributeV2.WinCmdline
			properties["win_script"] = scriptERIDiscoveryAttributeV2.WinScript
			d = append(d, &properties)
		}
	}
	return
}

func ScriptERIDiscoveryAttributeV2Model(d map[string]interface{}) *models.ScriptERIDiscoveryAttributeV2 {
	// assume that the incoming map only contains the relevant resource data
	groovyScript := d["groovy_script"].(string)
	linuxCmdline := d["linux_cmdline"].(string)
	linuxScript := d["linux_script"].(string)
	name := d["name"].(string)
	typeVar := d["type"].(string)
	winCmdline := d["win_cmdline"].(string)
	winScript := d["win_script"].(string)
	
	return &models.ScriptERIDiscoveryAttributeV2 {
		GroovyScript: groovyScript,
		LinuxCmdline: linuxCmdline,
		LinuxScript: linuxScript,
		Name: &name,
		Type: typeVar,
		WinCmdline: winCmdline,
		WinScript: winScript,
	}
}

func GetScriptERIDiscoveryAttributeV2PropertyFields() (t []string) {
	return []string{
		"groovy_script",
		"linux_cmdline",
		"linux_script",
		"name",
		"type",
		"win_cmdline",
		"win_script",
	}
}