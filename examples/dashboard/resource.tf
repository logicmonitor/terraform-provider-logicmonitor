resource "logicmonitor_dashboard" "mydashboard" {
	description = "my dashboard"
	name = "test_dashboard"
	sharable = true
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