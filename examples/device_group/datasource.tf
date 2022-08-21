resource "logicmonitor_device_group" "my_device_group" {
	description = "normal device group test" 
	disable_alerting = true
	enable_netflow = false
	default_collector_id = 10
	name = "meng test device group test"
	parent_id = 1
	custom_properties { 
          name = "addr"      
          value = "127.0.0.1" 
     }
	group_type = "Normal"
}


data "logicmonitor_device_group" "my_device_group" {
    filter = "description~\"normal device group test\""
	depends_on = [
		logicmonitor_device_group.my_device_group
	]
}

output "deviceGroup" {
  description = "devices group list"
  value       = data.logicmonitor_device_group.my_device_group.id
}