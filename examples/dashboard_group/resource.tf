resource "logicmonitor_dashboard_group" "myDashboardGroup"{
  name = "LogicMonitor Dashboard Group"
  description  = "LM dashboard group testing"
  widget_tokens = [
    {
      name  = "defaultDeviceGroup"
      value = "*"
      inherit_list = null
      type = null
    },
    {
      name  = "defaultServiceGroup"
      value = "*"
      inherit_list = null
      type = null
    }
  ]
}