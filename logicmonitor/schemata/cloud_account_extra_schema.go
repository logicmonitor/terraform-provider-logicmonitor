package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudAccountExtraSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account": {
			Type: schema.TypeList, //GoType: CloudAccount
			Elem: &schema.Resource{
				Schema: CloudAccountSchema(),
			},
			Required: true,
		},
		
		"default": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Required: true,
		},
		
		"devices": {
			Type: schema.TypeList, //GoType: CloudDevice
			Elem: &schema.Resource{
				Schema: CloudDeviceSchema(),
			},
			Computed: true,
		},
		
		"services": {
			Type: schema.TypeList, //GoType: CloudServices
			Elem: &schema.Resource{
				Schema: CloudServicesSchema(),
			},
			Required: true,
		},
		
	}
}

func SetCloudAccountExtraSubResourceData(m []*models.CloudAccountExtra) (d []*map[string]interface{}) {
	for _, cloudAccountExtra := range m {
		if cloudAccountExtra != nil {
			properties := make(map[string]interface{})
			properties["account"] = SetCloudAccountSubResourceData([]*models.CloudAccount{cloudAccountExtra.Account})
			properties["default"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudAccountExtra.Default})
			properties["devices"] = SetCloudDeviceSubResourceData([]*models.CloudDevice{cloudAccountExtra.Devices})
			properties["services"] = SetCloudServicesSubResourceData([]*models.CloudServices{cloudAccountExtra.Services})
			d = append(d, &properties)
		}
	}
	return
}

func CloudAccountExtraModel(d map[string]interface{}) *models.CloudAccountExtra {
	// assume that the incoming map only contains the relevant resource data
	var account *models.CloudAccount = nil
	accountList := d["account"].([]interface{})
	if len(accountList) > 0 { // len(nil) = 0
		account = CloudAccountModel(accountList[0].(map[string]interface{}))
	}
	var defaultVar *models.CloudServiceSettings = nil
	defaultList := d["default"].([]interface{})
	if len(defaultList) > 0 { // len(nil) = 0
		defaultVar = CloudServiceSettingsModel(defaultList[0].(map[string]interface{}))
	}
	var services *models.CloudServices = nil
	servicesList := d["services"].([]interface{})
	if len(servicesList) > 0 { // len(nil) = 0
		services = CloudServicesModel(servicesList[0].(map[string]interface{}))
	}
	
	return &models.CloudAccountExtra {
		Account: account,
		Default: defaultVar,
		Services: services,
	}
}

func GetCloudAccountExtraPropertyFields() (t []string) {
	return []string{
		"account",
		"default",
		"services",
	}
}