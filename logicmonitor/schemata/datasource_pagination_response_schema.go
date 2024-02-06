package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourcePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Datasource  
			Elem: &schema.Resource{
				Schema: DatasourceSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"search_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"total": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
	}
}

func SetDatasourcePaginationResponseSubResourceData(m []*models.DatasourcePaginationResponse) (d []*map[string]interface{}) {
	for _, datasourcePaginationResponse := range m {
		if datasourcePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDatasourceSubResourceData(datasourcePaginationResponse.Items)
			properties["search_id"] = datasourcePaginationResponse.SearchID
			properties["total"] = datasourcePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DatasourcePaginationResponseModel(d map[string]interface{}) *models.DatasourcePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Datasource)
	
	return &models.DatasourcePaginationResponse {
		Items: items,
	}
}

func GetDatasourcePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}