resource "logicmonitor_datasource" "my_datasource"{
  collect_interval = 100
  applies_to = data.logicmonitor_device.my_devices.id
  description = "test"
  collect_method = "aggregate"
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
  collector_attribute = [{
  name = "aggregate"
  }] 
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