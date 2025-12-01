package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ReportGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"matched_report_count": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"reports_count": {
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
// Only difference between this and ReportGroupSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceReportGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"matched_report_count": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"reports_count": {
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

func SetReportGroupResourceData(d *schema.ResourceData, m *models.ReportGroup) {
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("matched_report_count", m.MatchedReportCount)
	d.Set("name", m.Name)
	d.Set("reports_count", m.ReportsCount)
	d.Set("user_permission", m.UserPermission)
}

func SetReportGroupSubResourceData(m []*models.ReportGroup) (d []*map[string]interface{}) {
	for _, reportGroup := range m {
		if reportGroup != nil {
			properties := make(map[string]interface{})
			properties["description"] = reportGroup.Description
			properties["id"] = reportGroup.ID
			properties["matched_report_count"] = reportGroup.MatchedReportCount
			properties["name"] = reportGroup.Name
			properties["reports_count"] = reportGroup.ReportsCount
			properties["user_permission"] = reportGroup.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func ReportGroupModel(d *schema.ResourceData) *models.ReportGroup {
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	
	return &models.ReportGroup {
		Description: description,
		ID: int32(id),
		Name: &name,
	}
}
func GetReportGroupPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"name",
	}
}