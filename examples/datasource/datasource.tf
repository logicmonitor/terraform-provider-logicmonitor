resource "logicmonitor_datasource" "my_datasource"{
  collect_interval = 100
  has_multi_instances = true
  applies_to = "system.deviceId == \"22\""
  description = "test"
  collect_method = "script"
  eri_discovery_interval = 15
  enable_auto_discovery = true
  enable_eri_discovery = true
  eri_discovery_config {
    name = "ad_script"
    win_script = "string"
    groovy_script = "string"
    type = "string"
    linux_cmdline = "string"
    linux_script = "string"
    win_cmdline = "string"
  }
  name = "Amazon Website test"
  auto_discovery_config {
    persistent_instance = false
    schedule_interval = 0
    delete_inactive_instance = true
    method {
    name = "ad_script"
    }
    instance_auto_group_method = "none"
    instance_auto_group_method_params = ""
    filters = [
    {
     comment = "test"
     value = "test"
     operation = "string"
     attribute = "string"
    }]
    disable_instance = true
  }
  data_points = [{
          name = "CallCountTotal_mean8"
          description = "test"
          alert_for_no_data = 1
          max_digits = 0
          alert_clear_transition_interval = 0
          alert_transition_interval = 0
          data_type = 7
          max_value = ""
          min_value = ""
          alert_body = "string"
          alert_subject = "string"
          alert_expr = "string"
          alert_expr_note = "string"
          type = 2
          raw_data_field_name = "string"
          post_processor_method = "aggregation"
          post_processor_param = "{\"version\":\"1.0\",\"expression\":{\"funcName\":\"mean\",\"dataSourceName\":\"AWS_Cognito_GlobalAPICallStats\",\"dataPointName\":\"CallCountTotal\"},\"dataLack\":\"ignore\"}"
      }]
  display_name = "Testdemo"
  collector_attribute {
  name = "script"
  }
}

data "logicmonitor_datasource" "my_datasource" {
   filter = "description~\"datasource test\""
 	depends_on = [
		logicmonitor_datasource.my_datasource
 	]
}

output "datasource" {
  description = "datasource"
  value       = data.logicmonitor_datasource.my_datasource.id
}