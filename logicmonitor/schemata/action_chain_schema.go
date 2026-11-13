package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ActionChainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		
		"stages": {
			Type: schema.TypeList, //GoType: []*Stage
			Elem: &schema.Resource{
				Schema: StageSchema(),
			},
			Required: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and ActionChainSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceActionChainSchema() map[string]*schema.Schema {
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
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"stages": {
			Type: schema.TypeList, //GoType: []*Stage
			Elem: &schema.Resource{
				Schema: StageSchema(),
			},
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

// DataSourceActionChainListSchema is the schema of the plural "logicmonitor_action_chains" data source.
// "items" holds one read-only entry per ActionChain matching "filter", with the same attributes as DataSourceActionChainSchema().
func DataSourceActionChainListSchema() map[string]*schema.Schema {
	itemSchema := DataSourceActionChainSchema()
	delete(itemSchema, "filter")

	return map[string]*schema.Schema{
		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"page_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"total": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"items": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: ComputedOnlySchema(itemSchema)},
		},
	}
}

func SetActionChainResourceData(d *schema.ResourceData, m *models.ActionChain) {
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("stages", SetStageSubResourceData(m.Stages))
}

func SetActionChainSubResourceData(m []*models.ActionChain) (d []*map[string]interface{}) {
	for _, actionChain := range m {
		if actionChain != nil {
			properties := make(map[string]interface{})
			properties["description"] = actionChain.Description
			properties["id"] = actionChain.ID
			properties["name"] = actionChain.Name
			properties["stages"] = SetStageSubResourceData(actionChain.Stages)
			d = append(d, &properties)
		}
	}
	return
}

func ActionChainModel(d *schema.ResourceData) *models.ActionChain {
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	stagesTempSlice := d.Get("stages").([]interface{})
	stages := []*models.Stage{}
	for _, val := range stagesTempSlice {
		if m, ok := val.(map[string]interface{}); ok {
			stages = append(stages, StageModel(m))
		}
	}
	
	return &models.ActionChain {
		Description: description,
		ID: int32(id),
		Name: &name,
		Stages: stages,
	}
}
func GetActionChainPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"name",
		"stages",
	}
}