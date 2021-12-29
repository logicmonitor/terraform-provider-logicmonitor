package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applies_to": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"aws_regions_info": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"aws_test_result": {
			Type: schema.TypeList, //GoType: AwsAccountTestResult
			Elem: &schema.Resource{
				Schema: AwsAccountTestResultSchema(),
			},
			Computed: true,
		},
		
		"aws_test_result_code": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"azure_regions_info": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"azure_test_result": {
			Type: schema.TypeList, //GoType: AzureAccountTestResult
			Elem: &schema.Resource{
				Schema: AzureAccountTestResultSchema(),
			},
			Computed: true,
		},
		
		"azure_test_result_code": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"custom_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"default_collector_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"default_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"effective_alert_enabled": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"enable_netflow": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"extra": {
			Type: schema.TypeList, //GoType: CloudAccountExtra
			Elem: &schema.Resource{
				Schema: CloudAccountExtraSchema(),
			},
			Optional: true,
		},
		
		"full_path": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"gcp_regions_info": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"gcp_test_result": {
			Type: schema.TypeList, //GoType: GcpAccountTestResult
			Elem: &schema.Resource{
				Schema: GcpAccountTestResultSchema(),
			},
			Computed: true,
		},
		
		"gcp_test_result_code": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"group_status": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"group_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"has_netflow_enabled_devices": {
			Type: schema.TypeBool,
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
		
		"num_of_a_w_s_devices": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_azure_devices": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_direct_devices": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_direct_sub_groups": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_gcp_devices": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"num_of_hosts": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"parent_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"sub_groups": {
			Type: schema.TypeList, //GoType: []*DeviceGroupData 
			Elem: &schema.Resource{
				Schema: DeviceGroupDataSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetDeviceGroupResourceData(d *schema.ResourceData, m *models.DeviceGroup) {
	d.Set("applies_to", m.AppliesTo)
	d.Set("aws_regions_info", m.AwsRegionsInfo)
	d.Set("aws_test_result", SetAwsAccountTestResultSubResourceData([]*models.AwsAccountTestResult{m.AwsTestResult}))
	d.Set("aws_test_result_code", m.AwsTestResultCode)
	d.Set("azure_regions_info", m.AzureRegionsInfo)
	d.Set("azure_test_result", SetAzureAccountTestResultSubResourceData([]*models.AzureAccountTestResult{m.AzureTestResult}))
	d.Set("azure_test_result_code", m.AzureTestResultCode)
	d.Set("created_on", m.CreatedOn)
	d.Set("custom_properties", SetNameAndValueSubResourceData(m.CustomProperties))
	d.Set("default_collector_description", m.DefaultCollectorDescription)
	d.Set("default_collector_id", m.DefaultCollectorID)
	d.Set("description", m.Description)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("effective_alert_enabled", m.EffectiveAlertEnabled)
	d.Set("enable_netflow", m.EnableNetflow)
	d.Set("extra", SetCloudAccountExtraSubResourceData([]*models.CloudAccountExtra{m.Extra}))
	d.Set("full_path", m.FullPath)
	d.Set("gcp_regions_info", m.GcpRegionsInfo)
	d.Set("gcp_test_result", SetGcpAccountTestResultSubResourceData([]*models.GcpAccountTestResult{m.GcpTestResult}))
	d.Set("gcp_test_result_code", m.GcpTestResultCode)
	d.Set("group_status", m.GroupStatus)
	d.Set("group_type", m.GroupType)
	d.Set("has_netflow_enabled_devices", m.HasNetflowEnabledDevices)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("num_of_a_w_s_devices", m.NumOfAWSDevices)
	d.Set("num_of_azure_devices", m.NumOfAzureDevices)
	d.Set("num_of_direct_devices", m.NumOfDirectDevices)
	d.Set("num_of_direct_sub_groups", m.NumOfDirectSubGroups)
	d.Set("num_of_gcp_devices", m.NumOfGcpDevices)
	d.Set("num_of_hosts", m.NumOfHosts)
	d.Set("parent_id", m.ParentID)
	d.Set("sub_groups", SetDeviceGroupDataSubResourceData(m.SubGroups))
	d.Set("user_permission", m.UserPermission)
}

func SetDeviceGroupSubResourceData(m []*models.DeviceGroup) (d []*map[string]interface{}) {
	for _, deviceGroup := range m {
		if deviceGroup != nil {
			properties := make(map[string]interface{})
			properties["applies_to"] = deviceGroup.AppliesTo
			properties["aws_regions_info"] = deviceGroup.AwsRegionsInfo
			properties["aws_test_result"] = SetAwsAccountTestResultSubResourceData([]*models.AwsAccountTestResult{deviceGroup.AwsTestResult})
			properties["aws_test_result_code"] = deviceGroup.AwsTestResultCode
			properties["azure_regions_info"] = deviceGroup.AzureRegionsInfo
			properties["azure_test_result"] = SetAzureAccountTestResultSubResourceData([]*models.AzureAccountTestResult{deviceGroup.AzureTestResult})
			properties["azure_test_result_code"] = deviceGroup.AzureTestResultCode
			properties["created_on"] = deviceGroup.CreatedOn
			properties["custom_properties"] = SetNameAndValueSubResourceData(deviceGroup.CustomProperties)
			properties["default_collector_description"] = deviceGroup.DefaultCollectorDescription
			properties["default_collector_id"] = deviceGroup.DefaultCollectorID
			properties["description"] = deviceGroup.Description
			properties["disable_alerting"] = deviceGroup.DisableAlerting
			properties["effective_alert_enabled"] = deviceGroup.EffectiveAlertEnabled
			properties["enable_netflow"] = deviceGroup.EnableNetflow
			properties["extra"] = SetCloudAccountExtraSubResourceData([]*models.CloudAccountExtra{deviceGroup.Extra})
			properties["full_path"] = deviceGroup.FullPath
			properties["gcp_regions_info"] = deviceGroup.GcpRegionsInfo
			properties["gcp_test_result"] = SetGcpAccountTestResultSubResourceData([]*models.GcpAccountTestResult{deviceGroup.GcpTestResult})
			properties["gcp_test_result_code"] = deviceGroup.GcpTestResultCode
			properties["group_status"] = deviceGroup.GroupStatus
			properties["group_type"] = deviceGroup.GroupType
			properties["has_netflow_enabled_devices"] = deviceGroup.HasNetflowEnabledDevices
			properties["id"] = deviceGroup.ID
			properties["name"] = deviceGroup.Name
			properties["num_of_a_w_s_devices"] = deviceGroup.NumOfAWSDevices
			properties["num_of_azure_devices"] = deviceGroup.NumOfAzureDevices
			properties["num_of_direct_devices"] = deviceGroup.NumOfDirectDevices
			properties["num_of_direct_sub_groups"] = deviceGroup.NumOfDirectSubGroups
			properties["num_of_gcp_devices"] = deviceGroup.NumOfGcpDevices
			properties["num_of_hosts"] = deviceGroup.NumOfHosts
			properties["parent_id"] = deviceGroup.ParentID
			properties["sub_groups"] = SetDeviceGroupDataSubResourceData(deviceGroup.SubGroups)
			properties["user_permission"] = deviceGroup.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func DeviceGroupModel(d *schema.ResourceData) *models.DeviceGroup {
	appliesTo := d.Get("applies_to").(string)
	customProperties := utils.GetPropertiesFromResource(d, "custom_properties")
	defaultCollectorID := int32(d.Get("default_collector_id").(int))
	description := d.Get("description").(string)
	disableAlerting := d.Get("disable_alerting").(bool)
	enableNetflow := d.Get("enable_netflow").(bool)
	var extra *models.CloudAccountExtra = nil
	extraInterface, extraIsSet := d.GetOk("extra")
	if extraIsSet {
		extraMap := extraInterface.([]interface{})[0].(map[string]interface{})
		extra = CloudAccountExtraModel(extraMap)
	}
	groupType := d.Get("group_type").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	parentID := int32(d.Get("parent_id").(int))
	
	return &models.DeviceGroup {
		AppliesTo: appliesTo,
		CustomProperties: customProperties,
		DefaultCollectorID: defaultCollectorID,
		Description: description,
		DisableAlerting: disableAlerting,
		EnableNetflow: enableNetflow,
		Extra: extra,
		GroupType: groupType,
		ID: int32(id),
		Name: &name,
		ParentID: parentID,
	}
}
func GetDeviceGroupPropertyFields() (t []string) {
	return []string{
		"applies_to",
		"custom_properties",
		"default_collector_id",
		"description",
		"disable_alerting",
		"enable_netflow",
		"extra",
		"group_type",
		"id",
		"name",
		"parent_id",
	}
}