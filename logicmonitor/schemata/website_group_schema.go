package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func WebsiteGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"disable_alerting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"full_path": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"has_websites_disabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"num_of_direct_sub_groups": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"num_of_direct_websites": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"num_of_websites": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"stop_monitoring": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"test_location": {
			Type: schema.TypeList, //GoType: WebsiteLocation
			Elem: &schema.Resource{
				Schema: WebsiteLocationSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and WebsiteGroupSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceWebsiteGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"disable_alerting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"full_path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"has_websites_disabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"num_of_direct_sub_groups": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"num_of_direct_websites": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"num_of_websites": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"stop_monitoring": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"test_location": {
			Type: schema.TypeList, //GoType: WebsiteLocation
			Elem: &schema.Resource{
				Schema: WebsiteLocationSchema(),
			},
			Optional: true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetWebsiteGroupResourceData(d *schema.ResourceData, m *models.WebsiteGroup) {
	d.Set("description", m.Description)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("full_path", m.FullPath)
	d.Set("has_websites_disabled", m.HasWebsitesDisabled)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("num_of_direct_sub_groups", m.NumOfDirectSubGroups)
	d.Set("num_of_direct_websites", m.NumOfDirectWebsites)
	d.Set("num_of_websites", m.NumOfWebsites)
	d.Set("parent_id", m.ParentID)
	d.Set("properties", SetNameAndValueSubResourceData(m.Properties))
	d.Set("stop_monitoring", m.StopMonitoring)
	d.Set("test_location", SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{m.TestLocation}))
	d.Set("user_permission", m.UserPermission)
}

func SetWebsiteGroupSubResourceData(m []*models.WebsiteGroup) (d []*map[string]interface{}) {
	for _, websiteGroup := range m {
		if websiteGroup != nil {
			properties := make(map[string]interface{})
			properties["description"] = websiteGroup.Description
			properties["disable_alerting"] = websiteGroup.DisableAlerting
			properties["full_path"] = websiteGroup.FullPath
			properties["has_websites_disabled"] = websiteGroup.HasWebsitesDisabled
			properties["id"] = websiteGroup.ID
			properties["name"] = websiteGroup.Name
			properties["num_of_direct_sub_groups"] = websiteGroup.NumOfDirectSubGroups
			properties["num_of_direct_websites"] = websiteGroup.NumOfDirectWebsites
			properties["num_of_websites"] = websiteGroup.NumOfWebsites
			properties["parent_id"] = websiteGroup.ParentID
			properties["properties"] = SetNameAndValueSubResourceData(websiteGroup.Properties)
			properties["stop_monitoring"] = websiteGroup.StopMonitoring
			properties["test_location"] = SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{websiteGroup.TestLocation})
			properties["user_permission"] = websiteGroup.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteGroupModel(d *schema.ResourceData) *models.WebsiteGroup {
	description := d.Get("description").(string)
	disableAlerting := d.Get("disable_alerting").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	parentID := int32(d.Get("parent_id").(int))
	properties := utils.GetPropertiesFromResource(d, "properties")
	stopMonitoring := d.Get("stop_monitoring").(bool)
	testLocation := utils.GetPropFromLocationMap(d, "test_location")

	return &models.WebsiteGroup{
		Description:     description,
		DisableAlerting: disableAlerting,
		ID:              int32(id),
		Name:            &name,
		ParentID:        parentID,
		Properties:      properties,
		StopMonitoring:  stopMonitoring,
		TestLocation:    testLocation,
	}
}
func GetWebsiteGroupPropertyFields() (t []string) {
	return []string{
		"description",
		"disable_alerting",
		"id",
		"name",
		"parent_id",
		"properties",
		"stop_monitoring",
		"test_location",
	}
}
