package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertFiltersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acked": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"anomaly": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"chain": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"cleared": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"dependency_role": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"dependency_routing_state": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"group": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"instance": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"is_escalation": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"is_historical_sdt": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"keyword": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"rule": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"sdted": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"severity": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"suppression_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetAlertFiltersSubResourceData(m []*models.AlertFilters) (d []*map[string]interface{}) {
	for _, alertFilters := range m {
		if alertFilters != nil {
			properties := make(map[string]interface{})
			properties["acked"] = alertFilters.Acked
			properties["anomaly"] = alertFilters.Anomaly
			properties["chain"] = alertFilters.Chain
			properties["cleared"] = alertFilters.Cleared
			properties["data_point"] = alertFilters.DataPoint
			properties["data_source"] = alertFilters.DataSource
			properties["dependency_role"] = alertFilters.DependencyRole
			properties["dependency_routing_state"] = alertFilters.DependencyRoutingState
			properties["group"] = alertFilters.Group
			properties["host"] = alertFilters.Host
			properties["instance"] = alertFilters.Instance
			properties["is_escalation"] = alertFilters.IsEscalation
			properties["is_historical_sdt"] = alertFilters.IsHistoricalSdt
			properties["keyword"] = alertFilters.Keyword
			properties["rule"] = alertFilters.Rule
			properties["sdted"] = alertFilters.Sdted
			properties["severity"] = alertFilters.Severity
			properties["suppression_type"] = alertFilters.SuppressionType
			d = append(d, &properties)
		}
	}
	return
}

func AlertFiltersModel(d map[string]interface{}) *models.AlertFilters {
	// assume that the incoming map only contains the relevant resource data
	acked := d["acked"].(string)
	anomaly := d["anomaly"].(string)
	chain := d["chain"].(string)
	cleared := d["cleared"].(string)
	dataPoint := d["data_point"].(string)
	dataSource := d["data_source"].(string)
	dependencyRole := d["dependency_role"].(string)
	dependencyRoutingState := d["dependency_routing_state"].(string)
	group := d["group"].(string)
	host := d["host"].(string)
	instance := d["instance"].(string)
	isEscalation := d["is_escalation"].(string)
	isHistoricalSdt := d["is_historical_sdt"].(string)
	keyword := d["keyword"].(string)
	rule := d["rule"].(string)
	sdted := d["sdted"].(string)
	severity := d["severity"].(string)
	suppressionType := d["suppression_type"].(string)
	
	return &models.AlertFilters {
		Acked: acked,
		Anomaly: anomaly,
		Chain: chain,
		Cleared: cleared,
		DataPoint: dataPoint,
		DataSource: dataSource,
		DependencyRole: dependencyRole,
		DependencyRoutingState: dependencyRoutingState,
		Group: group,
		Host: host,
		Instance: instance,
		IsEscalation: isEscalation,
		IsHistoricalSdt: isHistoricalSdt,
		Keyword: keyword,
		Rule: rule,
		Sdted: sdted,
		Severity: severity,
		SuppressionType: suppressionType,
	}
}

func GetAlertFiltersPropertyFields() (t []string) {
	return []string{
		"acked",
		"anomaly",
		"chain",
		"cleared",
		"data_point",
		"data_source",
		"dependency_role",
		"dependency_routing_state",
		"group",
		"host",
		"instance",
		"is_escalation",
		"is_historical_sdt",
		"keyword",
		"rule",
		"sdted",
		"severity",
		"suppression_type",
	}
}