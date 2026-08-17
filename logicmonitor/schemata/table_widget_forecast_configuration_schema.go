package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetForecastConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"algorithm": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"confidence": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"severity": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"time_range": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetTableWidgetForecastConfigurationSubResourceData(m []*models.TableWidgetForecastConfiguration) (d []*map[string]interface{}) {
	for _, tableWidgetForecastConfiguration := range m {
		if tableWidgetForecastConfiguration != nil {
			properties := make(map[string]interface{})
			properties["algorithm"] = tableWidgetForecastConfiguration.Algorithm
			properties["confidence"] = tableWidgetForecastConfiguration.Confidence
			properties["severity"] = tableWidgetForecastConfiguration.Severity
			properties["time_range"] = tableWidgetForecastConfiguration.TimeRange
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetForecastConfigurationModel(d map[string]interface{}) *models.TableWidgetForecastConfiguration {
	// assume that the incoming map only contains the relevant resource data
	algorithm := d["algorithm"].(string)
	confidence := int32(d["confidence"].(int))
	severity := d["severity"].(string)
	timeRange := d["time_range"].(string)
	
	return &models.TableWidgetForecastConfiguration {
		Algorithm: algorithm,
		Confidence: confidence,
		Severity: severity,
		TimeRange: timeRange,
	}
}

func GetTableWidgetForecastConfigurationPropertyFields() (t []string) {
	return []string{
		"algorithm",
		"confidence",
		"severity",
		"time_range",
	}
}