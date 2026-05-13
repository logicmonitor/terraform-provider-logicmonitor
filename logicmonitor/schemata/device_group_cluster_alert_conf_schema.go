package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceGroupClusterAlertConfSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"count_by": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"data_point_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"data_source_display_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"device_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"min_alert_level": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"suppress_ind_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"threshold_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and DeviceGroupClusterAlertConfSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDeviceGroupClusterAlertConfSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"count_by": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_display_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"device_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"min_alert_level": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"suppress_ind_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"threshold_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetDeviceGroupClusterAlertConfResourceData(d *schema.ResourceData, m *models.DeviceGroupClusterAlertConf) {
	d.Set("alert_expr", m.AlertExpr)
	d.Set("count_by", m.CountBy)
	d.Set("data_point_description", m.DataPointDescription)
	d.Set("data_point_id", m.DataPointID)
	d.Set("data_point_name", m.DataPointName)
	d.Set("data_source_display_name", m.DataSourceDisplayName)
	d.Set("data_source_id", m.DataSourceID)
	d.Set("device_group_id", m.DeviceGroupID)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("min_alert_level", m.MinAlertLevel)
	d.Set("suppress_ind_alert", m.SuppressIndAlert)
	d.Set("threshold_type", m.ThresholdType)
}

func SetDeviceGroupClusterAlertConfSubResourceData(m []*models.DeviceGroupClusterAlertConf) (d []*map[string]interface{}) {
	for _, deviceGroupClusterAlertConf := range m {
		if deviceGroupClusterAlertConf != nil {
			properties := make(map[string]interface{})
			properties["alert_expr"] = deviceGroupClusterAlertConf.AlertExpr
			properties["count_by"] = deviceGroupClusterAlertConf.CountBy
			properties["data_point_description"] = deviceGroupClusterAlertConf.DataPointDescription
			properties["data_point_id"] = deviceGroupClusterAlertConf.DataPointID
			properties["data_point_name"] = deviceGroupClusterAlertConf.DataPointName
			properties["data_source_display_name"] = deviceGroupClusterAlertConf.DataSourceDisplayName
			properties["data_source_id"] = deviceGroupClusterAlertConf.DataSourceID
			properties["device_group_id"] = deviceGroupClusterAlertConf.DeviceGroupID
			properties["disable_alerting"] = deviceGroupClusterAlertConf.DisableAlerting
			properties["id"] = deviceGroupClusterAlertConf.ID
			properties["min_alert_level"] = deviceGroupClusterAlertConf.MinAlertLevel
			properties["suppress_ind_alert"] = deviceGroupClusterAlertConf.SuppressIndAlert
			properties["threshold_type"] = deviceGroupClusterAlertConf.ThresholdType
			d = append(d, &properties)
		}
	}
	return
}

func DeviceGroupClusterAlertConfModel(d *schema.ResourceData) *models.DeviceGroupClusterAlertConf {
	alertExpr := d.Get("alert_expr").(string)
	countBy := d.Get("count_by").(string)
	dataPointID := int32(d.Get("data_point_id").(int))
	dataSourceDisplayName := d.Get("data_source_display_name").(string)
	dataSourceID := int32(d.Get("data_source_id").(int))
	deviceGroupID := int32(d.Get("device_group_id").(int))
	disableAlerting := d.Get("disable_alerting").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	minAlertLevel := int32(d.Get("min_alert_level").(int))
	suppressIndAlert := d.Get("suppress_ind_alert").(bool)
	thresholdType := d.Get("threshold_type").(string)
	
	return &models.DeviceGroupClusterAlertConf {
		AlertExpr: alertExpr,
		CountBy: countBy,
		DataPointID: dataPointID,
		DataSourceDisplayName: dataSourceDisplayName,
		DataSourceID: dataSourceID,
		DeviceGroupID: deviceGroupID,
		DisableAlerting: disableAlerting,
		ID: int32(id),
		MinAlertLevel: minAlertLevel,
		SuppressIndAlert: suppressIndAlert,
		ThresholdType: thresholdType,
	}
}
func GetDeviceGroupClusterAlertConfPropertyFields() (t []string) {
	return []string{
		"alert_expr",
		"count_by",
		"data_point_id",
		"data_source_display_name",
		"data_source_id",
		"device_group_id",
		"disable_alerting",
		"id",
		"min_alert_level",
		"suppress_ind_alert",
		"threshold_type",
	}
}