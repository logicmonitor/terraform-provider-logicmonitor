resource "logicmonitor_device" "my_device"{
	auto_balanced_collector_group_id = 0
	//current_collector_id = 10
	custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		}
	]
	description = "This is a Cisco Router"
	device_type  = 0
	disable_alerting = true
	display_name = "Swagata  Cisco Router"
	enable_netflow = false
	link = "www.ciscorouter.com"
	name = "collector.host"
	netflow_collector_id = 1
	preferred_collector_id = 4
	related_device_id = -1
}


data "logicmonitor_device" "my_devices" {
    filter = "displayName~\"This is a Cisco Router\""
	depends_on = [
		logicmonitor_device.my_device
	]
}

output "devices" {
  description = "devices"
  value       = data.logicmonitor_device.my_devices.id
}