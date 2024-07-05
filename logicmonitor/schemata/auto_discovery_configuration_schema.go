package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AutoDiscoveryConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_source_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"delete_inactive_instance": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"disable_instance": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"filters": {
			Type: schema.TypeList, //GoType: []*AutoDiscoveryFilter  
			Elem: &schema.Resource{
				Schema: AutoDiscoveryFilterSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"instance_auto_group_method": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"instance_auto_group_method_params": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"method": {
			Type: schema.TypeList, //GoType: AutoDiscoveryMethod
			Elem: &schema.Resource{
				Schema: AutoDiscoveryMethodSchema(),
			},
			Required: true,
		},
		
		"persistent_instance": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"schedule_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
	}
}

func SetAutoDiscoveryConfigurationSubResourceData(m []*models.AutoDiscoveryConfiguration) (d []*map[string]interface{}) {
	for _, autoDiscoveryConfiguration := range m {
		if autoDiscoveryConfiguration != nil {
			properties := make(map[string]interface{})
			properties["data_source_name"] = autoDiscoveryConfiguration.DataSourceName
			properties["delete_inactive_instance"] = autoDiscoveryConfiguration.DeleteInactiveInstance
			properties["disable_instance"] = autoDiscoveryConfiguration.DisableInstance
			properties["filters"] = SetAutoDiscoveryFilterSubResourceData(autoDiscoveryConfiguration.Filters)
			properties["instance_auto_group_method"] = autoDiscoveryConfiguration.InstanceAutoGroupMethod
			properties["instance_auto_group_method_params"] = autoDiscoveryConfiguration.InstanceAutoGroupMethodParams
			properties["method"] = SetAutoDiscoveryMethodSubResourceData([]*models.AutoDiscoveryMethod{autoDiscoveryConfiguration.Method})
			properties["persistent_instance"] = autoDiscoveryConfiguration.PersistentInstance
			properties["schedule_interval"] = autoDiscoveryConfiguration.ScheduleInterval
			d = append(d, &properties)
		}
	}
	return
}

func AutoDiscoveryConfigurationModel(d map[string]interface{}) *models.AutoDiscoveryConfiguration {
	// assume that the incoming map only contains the relevant resource data
	deleteInactiveInstance := d["delete_inactive_instance"].(bool)
	disableInstance := d["disable_instance"].(bool)
	filters := utils.GetFilters(d["filters"].([]interface{}))
	instanceAutoGroupMethod := d["instance_auto_group_method"].(string)
	instanceAutoGroupMethodParams := d["instance_auto_group_method_params"].(string)
	var method *models.AutoDiscoveryMethod = nil
	methodList := d["method"].([]interface{})
	if len(methodList) > 0 { // len(nil) = 0
		method = AutoDiscoveryMethodModel(methodList[0].(map[string]interface{}))
	}
	persistentInstance := d["persistent_instance"].(bool)
	scheduleInterval := int32(d["schedule_interval"].(int))
	
	return &models.AutoDiscoveryConfiguration {
		DeleteInactiveInstance: deleteInactiveInstance,
		DisableInstance: disableInstance,
		Filters: filters,
		InstanceAutoGroupMethod: instanceAutoGroupMethod,
		InstanceAutoGroupMethodParams: instanceAutoGroupMethodParams,
		Method: method,
		PersistentInstance: persistentInstance,
		ScheduleInterval: scheduleInterval,
	}
}

func GetAutoDiscoveryConfigurationPropertyFields() (t []string) {
	return []string{
		"delete_inactive_instance",
		"disable_instance",
		"filters",
		"instance_auto_group_method",
		"instance_auto_group_method_params",
		"method",
		"persistent_instance",
		"schedule_interval",
	}
}