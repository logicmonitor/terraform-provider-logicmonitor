package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ChainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"period": {
			Type: schema.TypeList, //GoType: Period 
            Elem: &schema.Resource{
				Schema: PeriodSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"stages": {
			Type: schema.TypeList, //GoType: [][]*Recipient 
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: RecipientSchema()},
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetChainSubResourceData(m []*models.Chain) (d []*map[string]interface{}) {
	for _, chain := range m {
		if chain != nil {
			properties := make(map[string]interface{})
			properties["period"] = SetPeriodSubResourceData([]*models.Period{chain.Period})
			properties["stages"] = setChainStagesSubResourceData(chain.Stages)
			properties["type"] = chain.Type
			d = append(d, &properties)
		}
	}
	return
}
func setChainStagesSubResourceData(stages [][]*models.Recipient) []interface{} {
	if len(stages) == 0 {
		return []interface{}{}
	}

	result := make([]interface{}, 0, len(stages))
	for _, stage := range stages {
		if len(stage) == 0 {
			result = append(result, []interface{}{})
			continue
		}

		serializedStage := make([]interface{}, 0, len(stage))
		for _, recipient := range stage {
			if recipient == nil {
				continue
			}

			serializedStage = append(serializedStage, map[string]interface{}{
				"addr":    recipient.Addr,
				"contact": recipient.Contact,
				"method":  recipient.Method,
				"type":    recipient.Type,
			})
		}

		result = append(result, serializedStage)
	}

	return result
}

func ChainModel(d map[string]interface{}) *models.Chain {
	// assume that the incoming map only contains the relevant resource data
	var period *models.Period = nil
	periodList := d["period"].([]interface{})
	if len(periodList) > 0 { // len(nil) = 0
		period = PeriodModel(periodList[0].(map[string]interface{}))
	}
	stages := utils.GetRecipient(d["stages"].([]interface{}))
	typeVar := d["type"].(string)
	
	return &models.Chain {
		Period: period,
		Stages: stages,
		Type: &typeVar,
	}
}

func GetChainPropertyFields() (t []string) {
	return []string{
		"period",
		"stages",
		"type",
	}
}