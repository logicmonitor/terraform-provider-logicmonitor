resource "logicmonitor_dashboard" "mydashboard" {
	description = "my dashboard"
	name = "test_dashboard"
	sharable = true
}


data "logicmonitor_dashboard" "mydashboard"{
	filter = "description~\"my dashboard\""
	depends_on = [
		logicmonitor_dashboard.mydashboard
	]
}

output "dashboard" {
  description = "dashboard"
  value       = data.logicmonitor_dashboard.mydashboard
}