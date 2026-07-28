package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BigNumberInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"big_number_items": {
			Type: schema.TypeList, //GoType: []*BigNumberItem  
			Elem: &schema.Resource{
				Schema: BigNumberItemSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"counters": {
			Type: schema.TypeList, //GoType: []*Counter  
			Elem: &schema.Resource{
				Schema: CounterSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"data_points": {
			Type: schema.TypeList, //GoType: []*BigNumberDataPoint  
			Elem: &schema.Resource{
				Schema: BigNumberDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"virtual_data_points": {
			Type: schema.TypeList, //GoType: []*VirtualDataPoint  
			Elem: &schema.Resource{
				Schema: VirtualDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
	}
}

func SetBigNumberInfoSubResourceData(m []*models.BigNumberInfo) (d []*map[string]interface{}) {
	for _, bigNumberInfo := range m {
		if bigNumberInfo != nil {
			properties := make(map[string]interface{})
			properties["big_number_items"] = SetBigNumberItemSubResourceData(bigNumberInfo.BigNumberItems)
			properties["counters"] = SetCounterSubResourceData(bigNumberInfo.Counters)
			properties["data_points"] = SetBigNumberDataPointSubResourceData(bigNumberInfo.DataPoints)
			properties["virtual_data_points"] = SetVirtualDataPointSubResourceData(bigNumberInfo.VirtualDataPoints)
			d = append(d, &properties)
		}
	}
	return
}

func BigNumberInfoModel(d map[string]interface{}) *models.BigNumberInfo {
	// assume that the incoming map only contains the relevant resource data
	bigNumberItems := d["big_number_items"].([]*models.BigNumberItem)
	counters := d["counters"].([]*models.Counter)
	dataPoints := d["data_points"].([]*models.BigNumberDataPoint)
	virtualDataPoints := d["virtual_data_points"].([]*models.VirtualDataPoint)
	
	return &models.BigNumberInfo {
		BigNumberItems: bigNumberItems,
		Counters: counters,
		DataPoints: dataPoints,
		VirtualDataPoints: virtualDataPoints,
	}
}

func GetBigNumberInfoPropertyFields() (t []string) {
	return []string{
		"big_number_items",
		"counters",
		"data_points",
		"virtual_data_points",
	}
}