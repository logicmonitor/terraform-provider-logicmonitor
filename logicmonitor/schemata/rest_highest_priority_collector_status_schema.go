package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func RestHighestPriorityCollectorStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acked": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"in_s_d_t": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"is_down": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"status": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func SetRestHighestPriorityCollectorStatusSubResourceData(m []*models.RestHighestPriorityCollectorStatus) (d []*map[string]interface{}) {
	for _, restHighestPriorityCollectorStatus := range m {
		if restHighestPriorityCollectorStatus != nil {
			properties := make(map[string]interface{})
			properties["acked"] = restHighestPriorityCollectorStatus.Acked
			properties["in_s_d_t"] = restHighestPriorityCollectorStatus.InSDT
			properties["is_down"] = restHighestPriorityCollectorStatus.IsDown
			properties["status"] = restHighestPriorityCollectorStatus.Status
			d = append(d, &properties)
		}
	}
	return
}

func RestHighestPriorityCollectorStatusModel(d map[string]interface{}) *models.RestHighestPriorityCollectorStatus {
	// assume that the incoming map only contains the relevant resource data

	return &models.RestHighestPriorityCollectorStatus{}
}

func GetRestHighestPriorityCollectorStatusPropertyFields() (t []string) {
	return []string{}
}
