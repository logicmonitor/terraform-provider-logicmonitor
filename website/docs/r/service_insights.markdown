# Service_Insights by Terraform Provider LogicMonitor 

LM™ Service Insight is based on a new LogicMonitor resource type, Services. Services comprise of instances across one or more monitored devices. DataSources with a collection method of ‘Aggregate data’ are applied to Services, and the datapoints within these DataSources specify what data should be aggregated and how. Data is aggregated according to the DataSource across all instances that make up the Service.

## Steps to add services

- Add device resource from the provider with the device_type = 6. A service will be created.

   resource "logicmonitor_device" "my_device"{
	auto_balanced_collector_group_id = 0
	current_collector_id = 10
	custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		},
        {
            name = "predef.bizservice.evalMembersInterval",
            value = "30"
        },
        {
            name = "predef.bizservice.members",
            value = "{\"device\":[{\"deviceGroupFullPath\":\"Devices by Type*\",\"deviceDisplayName\":\"Test AWS Group TF1 Account\",\"deviceProperties\":[]}],\"instance\":[]}"
        }
	]
	description = "Test Service"
	device_type  = 6
	disable_alerting = true
	display_name = "Service test"
	enable_netflow = false
	link = "www.logicmonitor.com"
	name = "collector.host"
	netflow_collector_id = 1
	preferred_collector_id = 10
	related_device_id = -1
}

- After adding the service you need to configure service metrics and alert thresholds. For this you can add logicmonitor datasource by the provider with the datapoints configuration that should be aggregated up to the service level. 
  
- Use terraform datasource to get the deviceId for this service:

  data "logicmonitor_device" "my_devices" {
    filter = "displayName~\"Cisco Router Test\""
	depends_on = [
		logicmonitor_device.my_device
	]
}

output "devices" {
  description = "devices"
  value       = data.logicmonitor_device.my_devices.id
}

 - Use this deviceId in the applies_to field of logicmonitor datasource resource to configure the datasource for the service, e.g. applies_to = "system. deviceId == \"26\"" .

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


Here you can get more information about this resource : 
```sh
https://www.logicmonitor.com/support/lm-service-insight/adding-services
```

