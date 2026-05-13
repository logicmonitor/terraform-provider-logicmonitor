resource "logicmonitor_device_group_cluster_alert_conf" "my_device_group_cluster_alert_conf" {
  count_by = "host"
  data_source_id = 562268561
  min_alert_level = 3
  data_source_display_name = "TS_Alerts_PingTest"
  data_point_id = 323208
  disable_alerting = false
  suppress_ind_alert = true
  threshold_type = "absolute"
  alert_expr = "5 2 1"
  device_group_id = 253803
}

data "logicmonitor_device_group_cluster_alert_conf" "my_device_group_cluster_alert_conf" {
    filter = "threshold_type~\"\"absolute"
 	depends_on = [
		logicmonitor_device_group_cluster_alert_conf.my_device_group_cluster_alert_conf
 	]
}

output "ClusterAlertConf {
  description = "cluster alert conf"
  value       = data.logicmonitor_device_group_cluster_alert_conf.my_device_group_cluster_alert_conf.id
}

