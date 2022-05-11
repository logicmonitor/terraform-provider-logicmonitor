package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DashboardGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dashboards": {
			Type: schema.TypeList, //GoType: []*DashboardData 
			Elem: &schema.Resource{
				Schema: DashboardDataSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"full_path": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"num_of_dashboards": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_direct_dashboards": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_direct_sub_groups": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"parent_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"template": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"widget_tokens": {
			Type: schema.TypeList, //GoType: []*WidgetToken 
			Elem: &schema.Resource{
				Schema: WidgetTokenSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and DashboardGroupSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDashboardGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dashboards": {
			Type: schema.TypeList, //GoType: []*DashboardData 
			Elem: &schema.Resource{
				Schema: DashboardDataSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"full_path": {
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
		
		"num_of_dashboards": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"num_of_direct_dashboards": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"num_of_direct_sub_groups": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"parent_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"template": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"widget_tokens": {
			Type: schema.TypeList, //GoType: []*WidgetToken 
			Elem: &schema.Resource{
				Schema: WidgetTokenSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetDashboardGroupResourceData(d *schema.ResourceData, m *models.DashboardGroup) {
	d.Set("dashboards", SetDashboardDataSubResourceData(m.Dashboards))
	d.Set("description", m.Description)
	d.Set("full_path", m.FullPath)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("num_of_dashboards", m.NumOfDashboards)
	d.Set("num_of_direct_dashboards", m.NumOfDirectDashboards)
	d.Set("num_of_direct_sub_groups", m.NumOfDirectSubGroups)
	d.Set("parent_id", m.ParentID)
	d.Set("template", m.Template)
	d.Set("user_permission", m.UserPermission)
	d.Set("widget_tokens", SetWidgetTokenSubResourceData(m.WidgetTokens))
}

func SetDashboardGroupSubResourceData(m []*models.DashboardGroup) (d []*map[string]interface{}) {
	for _, dashboardGroup := range m {
		if dashboardGroup != nil {
			properties := make(map[string]interface{})
			properties["dashboards"] = SetDashboardDataSubResourceData(dashboardGroup.Dashboards)
			properties["description"] = dashboardGroup.Description
			properties["full_path"] = dashboardGroup.FullPath
			properties["id"] = dashboardGroup.ID
			properties["name"] = dashboardGroup.Name
			properties["num_of_dashboards"] = dashboardGroup.NumOfDashboards
			properties["num_of_direct_dashboards"] = dashboardGroup.NumOfDirectDashboards
			properties["num_of_direct_sub_groups"] = dashboardGroup.NumOfDirectSubGroups
			properties["parent_id"] = dashboardGroup.ParentID
			properties["template"] = dashboardGroup.Template
			properties["user_permission"] = dashboardGroup.UserPermission
			properties["widget_tokens"] = SetWidgetTokenSubResourceData(dashboardGroup.WidgetTokens)
			d = append(d, &properties)
		}
	}
	return
}

func DashboardGroupModel(d *schema.ResourceData) *models.DashboardGroup {
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	parentID := int32(d.Get("parent_id").(int))
	template := d.Get("template")
	widgetTokens := utils.GetPropFromWTMap(d, "widget_tokens")
	
	return &models.DashboardGroup {
		Description: description,
		ID: int32(id),
		Name: &name,
		ParentID: parentID,
		Template: template,
		WidgetTokens: widgetTokens,
	}
}
func GetDashboardGroupPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"name",
		"parent_id",
		"template",
		"widget_tokens",
	}
}