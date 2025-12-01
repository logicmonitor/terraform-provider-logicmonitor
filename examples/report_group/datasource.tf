resource "logicmonitor_report_group" "my_report_group" {
  name        = "TF Report Group"
  description = "report group test"
}
data "logicmonitor_report_group" "my_report_group" {
   filter = "description~\"report group test\""
 	depends_on = [
		logicmonitor_report_group.my_report_group
 	]
}

output "reportGroup" {
  description = "report group"
  value       = data.logicmonitor_report_group.my_report_group.id
}