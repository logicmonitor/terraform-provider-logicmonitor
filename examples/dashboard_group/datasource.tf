resource "logicmonitor_dashboard_group" "myDashboardGroup"{
  name = "LogicMonitor Dashboard Group"
  description  = "LM dashboard group testing"
}


data "logicmonitor_dashboard_group" "myDashboardGroup" {
    filter = "description~\"LM dashboard group testing\""
	depends_on = [
		logicmonitor_dashboard_group.myDashboardGroup
	]
}

output "dashboardGroups" {
  description = "dashboard groups"
  value       = data.logicmonitor_dashboard_group.myDashboardGroup
}