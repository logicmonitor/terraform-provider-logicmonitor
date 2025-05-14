package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_balanced_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"auto_props_assigned_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"auto_props_updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"aws_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"azure_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"collector_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"current_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"deleted_time_in_ms": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"display_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"enable_netflow": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"gcp_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"host_group_ids": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_status": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"inherited_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"last_data_time": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"last_rawdata_time": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"link": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"need_stc_grp_and_sorted_c_p": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netflow_collector_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"netflow_collector_group_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"netflow_collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"netflow_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"preferred_collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"preferred_collector_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"related_device_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"scan_config_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"system_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"to_delete_time_in_ms": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"up_time_in_seconds": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"updated_on": {
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
// Only difference between this and DeviceSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_balanced_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"auto_props_assigned_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_props_updated_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"aws_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"azure_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"collector_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"current_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"deleted_time_in_ms": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"display_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable_netflow": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"gcp_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"host_group_ids": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_status": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"inherited_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"last_data_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"last_rawdata_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"link": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"need_stc_grp_and_sorted_c_p": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netflow_collector_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"netflow_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"netflow_collector_group_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"netflow_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"preferred_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"related_device_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"scan_config_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"system_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"to_delete_time_in_ms": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"up_time_in_seconds": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"updated_on": {
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

func SetDeviceResourceData(d *schema.ResourceData, m *models.Device) {
	d.Set("auto_balanced_collector_group_id", m.AutoBalancedCollectorGroupID)
	d.Set("auto_properties", SetNameAndValueSubResourceData(m.AutoProperties))
	d.Set("auto_props_assigned_on", m.AutoPropsAssignedOn)
	d.Set("auto_props_updated_on", m.AutoPropsUpdatedOn)
	d.Set("aws_state", m.AwsState)
	d.Set("azure_state", m.AzureState)
	d.Set("collector_description", m.CollectorDescription)
	d.Set("created_on", m.CreatedOn)
	d.Set("current_collector_id", m.CurrentCollectorID)
	d.Set("custom_properties", SetNameAndValueSubResourceData(m.CustomProperties))
	d.Set("deleted_time_in_ms", m.DeletedTimeInMs)
	d.Set("description", m.Description)
	d.Set("device_type", m.DeviceType)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("display_name", m.DisplayName)
	d.Set("enable_netflow", m.EnableNetflow)
	d.Set("gcp_state", m.GcpState)
	d.Set("host_group_ids", m.HostGroupIds)
	d.Set("host_status", m.HostStatus)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("inherited_properties", SetNameAndValueSubResourceData(m.InheritedProperties))
	d.Set("last_data_time", m.LastDataTime)
	d.Set("last_rawdata_time", m.LastRawdataTime)
	d.Set("link", m.Link)
	d.Set("name", m.Name)
	// Special handling for 'need_stc_grp_and_sorted_c_p' as it is a query param and not returned by API
	if val, ok := d.GetOk("need_stc_grp_and_sorted_c_p"); ok {
		d.Set("need_stc_grp_and_sorted_c_p", val)
	} else {
		d.Set("need_stc_grp_and_sorted_c_p", m.NeedStcGrpAndSortedCP)
	}
	d.Set("netflow_collector_description", m.NetflowCollectorDescription)
	d.Set("netflow_collector_group_id", m.NetflowCollectorGroupID)
	d.Set("netflow_collector_group_name", m.NetflowCollectorGroupName)
	d.Set("netflow_collector_id", m.NetflowCollectorID)
	d.Set("preferred_collector_group_id", m.PreferredCollectorGroupID)
	d.Set("preferred_collector_group_name", m.PreferredCollectorGroupName)
	d.Set("preferred_collector_id", m.PreferredCollectorID)
	d.Set("related_device_id", m.RelatedDeviceID)
	d.Set("scan_config_id", m.ScanConfigID)
	d.Set("system_properties", SetNameAndValueSubResourceData(m.SystemProperties))
	d.Set("to_delete_time_in_ms", m.ToDeleteTimeInMs)
	d.Set("up_time_in_seconds", m.UpTimeInSeconds)
	d.Set("updated_on", m.UpdatedOn)
	d.Set("user_permission", m.UserPermission)
}

func SetDeviceSubResourceData(m []*models.Device) (d []*map[string]interface{}) {
	for _, device := range m {
		if device != nil {
			properties := make(map[string]interface{})
			properties["auto_balanced_collector_group_id"] = device.AutoBalancedCollectorGroupID
			properties["auto_properties"] = SetNameAndValueSubResourceData(device.AutoProperties)
			properties["auto_props_assigned_on"] = device.AutoPropsAssignedOn
			properties["auto_props_updated_on"] = device.AutoPropsUpdatedOn
			properties["aws_state"] = device.AwsState
			properties["azure_state"] = device.AzureState
			properties["collector_description"] = device.CollectorDescription
			properties["created_on"] = device.CreatedOn
			properties["current_collector_id"] = device.CurrentCollectorID
			properties["custom_properties"] = SetNameAndValueSubResourceData(device.CustomProperties)
			properties["deleted_time_in_ms"] = device.DeletedTimeInMs
			properties["description"] = device.Description
			properties["device_type"] = device.DeviceType
			properties["disable_alerting"] = device.DisableAlerting
			properties["display_name"] = device.DisplayName
			properties["enable_netflow"] = device.EnableNetflow
			properties["gcp_state"] = device.GcpState
			properties["host_group_ids"] = device.HostGroupIds
			properties["host_status"] = device.HostStatus
			properties["id"] = device.ID
			properties["inherited_properties"] = SetNameAndValueSubResourceData(device.InheritedProperties)
			properties["last_data_time"] = device.LastDataTime
			properties["last_rawdata_time"] = device.LastRawdataTime
			properties["link"] = device.Link
			properties["name"] = device.Name
			properties["need_stc_grp_and_sorted_c_p"] = device.NeedStcGrpAndSortedCP
			properties["netflow_collector_description"] = device.NetflowCollectorDescription
			properties["netflow_collector_group_id"] = device.NetflowCollectorGroupID
			properties["netflow_collector_group_name"] = device.NetflowCollectorGroupName
			properties["netflow_collector_id"] = device.NetflowCollectorID
			properties["preferred_collector_group_id"] = device.PreferredCollectorGroupID
			properties["preferred_collector_group_name"] = device.PreferredCollectorGroupName
			properties["preferred_collector_id"] = device.PreferredCollectorID
			properties["related_device_id"] = device.RelatedDeviceID
			properties["scan_config_id"] = device.ScanConfigID
			properties["system_properties"] = SetNameAndValueSubResourceData(device.SystemProperties)
			properties["to_delete_time_in_ms"] = device.ToDeleteTimeInMs
			properties["up_time_in_seconds"] = device.UpTimeInSeconds
			properties["updated_on"] = device.UpdatedOn
			properties["user_permission"] = device.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func DeviceModel(d *schema.ResourceData) *models.Device {
	autoBalancedCollectorGroupID := int32(d.Get("auto_balanced_collector_group_id").(int))
	currentCollectorID := int32(d.Get("current_collector_id").(int))
	customProperties := utils.GetPropertiesFromResource(d, "custom_properties")
	description := d.Get("description").(string)
	deviceType := int32(d.Get("device_type").(int))
	disableAlerting := d.Get("disable_alerting").(bool)
	displayName := d.Get("display_name").(string)
	enableNetflow := d.Get("enable_netflow").(bool)
	hostGroupIds := d.Get("host_group_ids").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	link := d.Get("link").(string)
	name := d.Get("name").(string)
	needStcGrpAndSortedCP := d.Get("need_stc_grp_and_sorted_c_p").(bool)
	netflowCollectorID := int32(d.Get("netflow_collector_id").(int))
	preferredCollectorID := int32(d.Get("preferred_collector_id").(int))
	relatedDeviceID := int32(d.Get("related_device_id").(int))
	
	return &models.Device {
		AutoBalancedCollectorGroupID: autoBalancedCollectorGroupID,
		CurrentCollectorID: currentCollectorID,
		CustomProperties: customProperties,
		Description: description,
		DeviceType: deviceType,
		DisableAlerting: disableAlerting,
		DisplayName: &displayName,
		EnableNetflow: enableNetflow,
		HostGroupIds: hostGroupIds,
		ID: int32(id),
		Link: link,
		Name: &name,
		NeedStcGrpAndSortedCP: &needStcGrpAndSortedCP,
		NetflowCollectorID: netflowCollectorID,
		PreferredCollectorID: &preferredCollectorID,
		RelatedDeviceID: relatedDeviceID,
	}
}
func GetDevicePropertyFields() (t []string) {
	return []string{
		"auto_balanced_collector_group_id",
		"current_collector_id",
		"custom_properties",
		"description",
		"device_type",
		"disable_alerting",
		"display_name",
		"enable_netflow",
		"host_group_ids",
		"id",
		"link",
		"name",
		"need_stc_grp_and_sorted_c_p",
		"netflow_collector_id",
		"preferred_collector_id",
		"related_device_id",
	}
}