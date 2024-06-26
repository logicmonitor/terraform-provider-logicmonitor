// For adding a logicmonitor datasource resource , also use this to configure a datasource for the service resource. For more information refer  README.md file (terraform-provider-logicmonitor/resources_docs/README.md) 

resource "logicmonitor_datasource" "my_datasource"{
  collect_interval = 100
  applies_to = "system.deviceId == \"22\""
  description = "test"
  collect_method = "script"
  eri_discovery_interval = 1
  enable_eri_discovery = true
  name = "Amazon Website test"
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