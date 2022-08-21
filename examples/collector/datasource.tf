resource "logicmonitor_collector" "my_collector"{
	arch = "linux64"
	automatic_upgrade_info {
		day_of_week = "MON"
		hour = 15
		timezone = "America/Los_Angeles"
		description = "regular MGD updates"
		occurrence = "Any"
		version = "MGD"
		minute = 0
	}
	build = "30.103"
	collector_group_id = 1
	collector_size = "nano"
	company = "lm"
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
	description = "Linux Collector"
	ea = false
	enable_fail_back = true
	enable_fail_over_on_collector_device = true
	escalating_chain_id = 1
	monitor_others = true
	need_auto_create_collector_device = true
	number_of_instances = 0
	resend_ival = 15
  	specified_collector_device_group_id = 0
  	suppress_alert_clear = true
}

data "logicmonitor_collector" "my_collector" {
    filter = "description~\"Linux Collector\""
 	depends_on = [
		logicmonitor_collector.my_collector
 	]
}

output "collector" {
  description = "collector"
  value       = data.logicmonitor_collector.my_collector.id
}