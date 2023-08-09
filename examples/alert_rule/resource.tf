resource "logicmonitor_alert_rule" "my_alert_rule"{
  datapoint = "*"
  instance = "*"
  devices = [
    "Cisco Router"
  ]
  escalating_chain_id = 9
  priority = 100
  suppress_alert_ack_sdt = true
  datasource = "Port-"
  suppress_alert_clear = true
  name = "Alert Rule Test"
  level_str = "Critical"
  device_groups = [
    "Devices by Type"
  ]
  escalation_interval = 15
}