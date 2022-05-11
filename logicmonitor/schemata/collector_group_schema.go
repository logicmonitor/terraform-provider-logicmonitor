package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CollectorGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_balance": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"auto_balance_instance_count_threshold": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_balance_strategy": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"create_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"custom_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"num_of_collectors": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and CollectorGroupSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceCollectorGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_balance": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"auto_balance_instance_count_threshold": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_balance_strategy": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"create_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"num_of_collectors": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetCollectorGroupResourceData(d *schema.ResourceData, m *models.CollectorGroup) {
	d.Set("auto_balance", m.AutoBalance)
	d.Set("auto_balance_instance_count_threshold", m.AutoBalanceInstanceCountThreshold)
	d.Set("auto_balance_strategy", m.AutoBalanceStrategy)
	d.Set("create_on", m.CreateOn)
	d.Set("custom_properties", SetNameAndValueSubResourceData(m.CustomProperties))
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("num_of_collectors", m.NumOfCollectors)
	d.Set("user_permission", m.UserPermission)
}

func SetCollectorGroupSubResourceData(m []*models.CollectorGroup) (d []*map[string]interface{}) {
	for _, collectorGroup := range m {
		if collectorGroup != nil {
			properties := make(map[string]interface{})
			properties["auto_balance"] = collectorGroup.AutoBalance
			properties["auto_balance_instance_count_threshold"] = collectorGroup.AutoBalanceInstanceCountThreshold
			properties["auto_balance_strategy"] = collectorGroup.AutoBalanceStrategy
			properties["create_on"] = collectorGroup.CreateOn
			properties["custom_properties"] = SetNameAndValueSubResourceData(collectorGroup.CustomProperties)
			properties["description"] = collectorGroup.Description
			properties["id"] = collectorGroup.ID
			properties["name"] = collectorGroup.Name
			properties["num_of_collectors"] = collectorGroup.NumOfCollectors
			properties["user_permission"] = collectorGroup.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func CollectorGroupModel(d *schema.ResourceData) *models.CollectorGroup {
	autoBalance := d.Get("auto_balance").(bool)
	autoBalanceInstanceCountThreshold := int32(d.Get("auto_balance_instance_count_threshold").(int))
	autoBalanceStrategy := d.Get("auto_balance_strategy").(string)
	customProperties := utils.GetPropertiesFromResource(d, "custom_properties")
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	
	return &models.CollectorGroup {
		AutoBalance: autoBalance,
		AutoBalanceInstanceCountThreshold: autoBalanceInstanceCountThreshold,
		AutoBalanceStrategy: autoBalanceStrategy,
		CustomProperties: customProperties,
		Description: description,
		ID: int32(id),
		Name: &name,
	}
}
func GetCollectorGroupPropertyFields() (t []string) {
	return []string{
		"auto_balance",
		"auto_balance_instance_count_threshold",
		"auto_balance_strategy",
		"custom_properties",
		"description",
		"id",
		"name",
	}
}