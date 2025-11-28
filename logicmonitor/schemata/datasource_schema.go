package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func DatasourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_group_ids": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"access_groups": {
			Type: schema.TypeList, //GoType: []*AccessGroup
			Elem: &schema.Resource{
				Schema: AccessGroupSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"applies_to": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"audit_version": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"auto_discovery_config": {
			Type: schema.TypeList, //GoType: AutoDiscoveryConfiguration
			Elem: &schema.Resource{
				Schema: AutoDiscoveryConfigurationSchema(),
			},
			Optional: true,
		},

		"checksum": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"collect_interval": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"collect_method": {
			Type:     schema.TypeString,
			Required: true,
		},

		"collector_attribute": {
			Type: schema.TypeList, //GoType: CollectorAttribute
			Elem: &schema.Resource{
				Schema: CollectorAttributeSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"data_points": {
			Type: schema.TypeList, //GoType: []*DataPoint
			Elem: &schema.Resource{
				Schema: DataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enable_auto_discovery": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"enable_eri_discovery": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"eri_discovery_config": {
			Type: schema.TypeList, //GoType: ScriptERIDiscoveryAttributeV2
			Elem: &schema.Resource{
				Schema: ScriptERIDiscoveryAttributeV2Schema(),
			},
			Optional: true,
		},

		"eri_discovery_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"has_multi_instances": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"installation_metadata": {
			Type: schema.TypeList, //GoType: IntegrationMetadata
			Elem: &schema.Resource{
				Schema: IntegrationMetadataSchema(),
			},
			Computed: true,
		},

		"lineage_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"payload_version": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"tags": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"technology": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"use_wild_value_as_uuid": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and DatasourceSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDatasourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_group_ids": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"access_groups": {
			Type: schema.TypeList, //GoType: []*AccessGroup
			Elem: &schema.Resource{
				Schema: AccessGroupSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"applies_to": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"audit_version": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"auto_discovery_config": {
			Type: schema.TypeList, //GoType: AutoDiscoveryConfiguration
			Elem: &schema.Resource{
				Schema: AutoDiscoveryConfigurationSchema(),
			},
			Optional: true,
		},

		"checksum": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"collect_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"collect_method": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"collector_attribute": {
			Type: schema.TypeList, //GoType: CollectorAttribute
			Elem: &schema.Resource{
				Schema: CollectorAttributeSchema(),
			},
			Optional: true,
		},

		"data_points": {
			Type: schema.TypeList, //GoType: []*DataPoint
			Elem: &schema.Resource{
				Schema: DataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enable_auto_discovery": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"enable_eri_discovery": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"eri_discovery_config": {
			Type: schema.TypeList, //GoType: ScriptERIDiscoveryAttributeV2
			Elem: &schema.Resource{
				Schema: ScriptERIDiscoveryAttributeV2Schema(),
			},
			Optional: true,
		},

		"eri_discovery_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"has_multi_instances": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"installation_metadata": {
			Type: schema.TypeList, //GoType: IntegrationMetadata
			Elem: &schema.Resource{
				Schema: IntegrationMetadataSchema(),
			},
			Optional: true,
		},

		"lineage_id": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"payload_version": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"tags": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"technology": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"use_wild_value_as_uuid": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"version": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetDatasourceResourceData(d *schema.ResourceData, m *models.Datasource) {
	d.Set("access_group_ids", m.AccessGroupIds)
	d.Set("access_groups", SetAccessGroupSubResourceData(m.AccessGroups))
	d.Set("applies_to", m.AppliesTo)
	d.Set("audit_version", m.AuditVersion)
	d.Set("auto_discovery_config", SetAutoDiscoveryConfigurationSubResourceData([]*models.AutoDiscoveryConfiguration{m.AutoDiscoveryConfig}))
	d.Set("checksum", m.Checksum)
	d.Set("collect_interval", m.CollectInterval)
	d.Set("collect_method", m.CollectMethod)
	d.Set("collector_attribute", SetCollectorAttributeSubResourceData([]*models.CollectorAttribute{m.CollectorAttribute}))
	d.Set("data_points", SetDataPointSubResourceData(m.DataPoints))
	d.Set("description", m.Description)
	d.Set("display_name", m.DisplayName)
	d.Set("enable_auto_discovery", m.EnableAutoDiscovery)
	d.Set("enable_eri_discovery", m.EnableEriDiscovery)
	d.Set("eri_discovery_config", SetScriptERIDiscoveryAttributeV2SubResourceData([]*models.ScriptERIDiscoveryAttributeV2{m.EriDiscoveryConfig}))
	d.Set("eri_discovery_interval", m.EriDiscoveryInterval)
	d.Set("group", m.Group)
	d.Set("has_multi_instances", m.HasMultiInstances)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("installation_metadata", SetIntegrationMetadataSubResourceData([]*models.IntegrationMetadata{m.InstallationMetadata}))
	d.Set("lineage_id", m.LineageID)
	d.Set("name", m.Name)
	d.Set("payload_version", m.PayloadVersion)
	d.Set("tags", m.Tags)
	d.Set("technology", m.Technology)
	d.Set("use_wild_value_as_uuid", m.UseWildValueAsUUID)
	d.Set("version", m.Version)
}

func SetDatasourceSubResourceData(m []*models.Datasource) (d []*map[string]interface{}) {
	for _, datasource := range m {
		if datasource != nil {
			properties := make(map[string]interface{})
			properties["access_group_ids"] = datasource.AccessGroupIds
			properties["access_groups"] = SetAccessGroupSubResourceData(datasource.AccessGroups)
			properties["applies_to"] = datasource.AppliesTo
			properties["audit_version"] = datasource.AuditVersion
			properties["auto_discovery_config"] = SetAutoDiscoveryConfigurationSubResourceData([]*models.AutoDiscoveryConfiguration{datasource.AutoDiscoveryConfig})
			properties["checksum"] = datasource.Checksum
			properties["collect_interval"] = datasource.CollectInterval
			properties["collect_method"] = datasource.CollectMethod
			properties["collector_attribute"] = SetCollectorAttributeSubResourceData([]*models.CollectorAttribute{datasource.CollectorAttribute})
			properties["data_points"] = SetDataPointSubResourceData(datasource.DataPoints)
			properties["description"] = datasource.Description
			properties["display_name"] = datasource.DisplayName
			properties["enable_auto_discovery"] = datasource.EnableAutoDiscovery
			properties["enable_eri_discovery"] = datasource.EnableEriDiscovery
			properties["eri_discovery_config"] = SetScriptERIDiscoveryAttributeV2SubResourceData([]*models.ScriptERIDiscoveryAttributeV2{datasource.EriDiscoveryConfig})
			properties["eri_discovery_interval"] = datasource.EriDiscoveryInterval
			properties["group"] = datasource.Group
			properties["has_multi_instances"] = datasource.HasMultiInstances
			properties["id"] = datasource.ID
			properties["installation_metadata"] = SetIntegrationMetadataSubResourceData([]*models.IntegrationMetadata{datasource.InstallationMetadata})
			properties["lineage_id"] = datasource.LineageID
			properties["name"] = datasource.Name
			properties["payload_version"] = datasource.PayloadVersion
			properties["tags"] = datasource.Tags
			properties["technology"] = datasource.Technology
			properties["use_wild_value_as_uuid"] = datasource.UseWildValueAsUUID
			properties["version"] = datasource.Version
			d = append(d, &properties)
		}
	}
	return
}

func DatasourceModel(d *schema.ResourceData) *models.Datasource {
	accessGroupIds := utils.ConvertSetToInt32Slice(d.Get("access_group_ids").([]interface{}))
	accessGroupsTempSlice := d.Get("access_groups").([]interface{})
	accessGroups := []*models.AccessGroup{}
	for _, val := range accessGroupsTempSlice {
		if m, ok := val.(map[string]interface{}); ok {
			accessGroups = append(accessGroups, AccessGroupModel(m))
		}
	}
	appliesTo := d.Get("applies_to").(string)
	var autoDiscoveryConfig *models.AutoDiscoveryConfiguration = nil
	autoDiscoveryConfigInterface, autoDiscoveryConfigIsSet := d.GetOk("auto_discovery_config")
	if autoDiscoveryConfigIsSet {
		autoDiscoveryConfigMap := autoDiscoveryConfigInterface.([]interface{})[0].(map[string]interface{})
		autoDiscoveryConfig = AutoDiscoveryConfigurationModel(autoDiscoveryConfigMap)
	}
	collectInterval := int32(d.Get("collect_interval").(int))
	collectMethod := d.Get("collect_method").(string)
	var collectorAttribute *models.CollectorAttribute = nil
	collectorAttributeInterface, collectorAttributeIsSet := d.GetOk("collector_attribute")
	if collectorAttributeIsSet {
		collectorAttributeMap := collectorAttributeInterface.([]interface{})[0].(map[string]interface{})
		collectorAttribute = CollectorAttributeModel(collectorAttributeMap)
	}
	dataPoints := utils.GetPropFromDatapoint(d, "data_points")
	description := d.Get("description").(string)
	displayName := d.Get("display_name").(string)
	enableAutoDiscovery := d.Get("enable_auto_discovery").(bool)
	enableEriDiscovery := d.Get("enable_eri_discovery").(bool)
	var eriDiscoveryConfig *models.ScriptERIDiscoveryAttributeV2 = nil
	eriDiscoveryConfigInterface, eriDiscoveryConfigIsSet := d.GetOk("eri_discovery_config")
	if eriDiscoveryConfigIsSet {
		eriDiscoveryConfigMap := eriDiscoveryConfigInterface.([]interface{})[0].(map[string]interface{})
		eriDiscoveryConfig = ScriptERIDiscoveryAttributeV2Model(eriDiscoveryConfigMap)
	}
	eriDiscoveryInterval := int32(d.Get("eri_discovery_interval").(int))
	group := d.Get("group").(string)
	hasMultiInstances := d.Get("has_multi_instances").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	tags := d.Get("tags").(string)
	technology := d.Get("technology").(string)

	return &models.Datasource{
		AccessGroupIds:       accessGroupIds,
		AccessGroups:         accessGroups,
		AppliesTo:            appliesTo,
		AutoDiscoveryConfig:  autoDiscoveryConfig,
		CollectInterval:      &collectInterval,
		CollectMethod:        &collectMethod,
		CollectorAttribute:   collectorAttribute,
		DataPoints:           dataPoints,
		Description:          description,
		DisplayName:          displayName,
		EnableAutoDiscovery:  enableAutoDiscovery,
		EnableEriDiscovery:   enableEriDiscovery,
		EriDiscoveryConfig:   eriDiscoveryConfig,
		EriDiscoveryInterval: eriDiscoveryInterval,
		Group:                group,
		HasMultiInstances:    hasMultiInstances,
		ID:                   int32(id),
		Name:                 &name,
		Tags:                 tags,
		Technology:           technology,
	}
}
func GetDatasourcePropertyFields() (t []string) {
	return []string{
		"access_group_ids",
		"access_groups",
		"applies_to",
		"auto_discovery_config",
		"collect_interval",
		"collect_method",
		"collector_attribute",
		"data_points",
		"description",
		"display_name",
		"enable_auto_discovery",
		"enable_eri_discovery",
		"eri_discovery_config",
		"eri_discovery_interval",
		"group",
		"has_multi_instances",
		"id",
		"name",
		"tags",
		"technology",
	}
}
