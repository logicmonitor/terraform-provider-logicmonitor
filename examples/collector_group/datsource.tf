resource "logicmonitor_collector_group" "my_collector_group"{
    auto_balance = false
    auto_balance_instance_count_threshold = 0
    auto_balance_strategy = ""
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
    description = "Group for collectors dedicated to Network Devices"
    name = "Collector_Group_(Network Devices)"
}


data "logicmonitor_collector_group" "my_collector_group" {
    filter = "description~\"Group for collectors dedicated to Network Devices\""
	depends_on = [
 		logicmonitor_collector_group.my_collector_group
	]
}

output "collector_group" {
   description = "collector group list"
   value       = data.logicmonitor_collector_group.my_collector_group
}