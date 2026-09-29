resource "logicmonitor_dashboard" "mydashboard" {
	description = "my dashboard"
	name = "test_dashboard"
	sharable = true
  group_id = 1
  # Pass the entire dashboard import JSON as a single string. You can also use
  # template = file("${path.module}/dashboard-export.json") with a raw export.
  template = jsonencode({
    name                 = "test dashboard"
    description          = "Created from a dashboard template"
    widgetsConfigVersion = 2
    type                 = "dashboard"
    version              = 2
    widgetTokens = [
      {
        name  = "defaultDeviceGroup"
        value = "*"
      },
      {
        name  = "defaultServiceGroup"
        value = "*"
      }
    ]
    widgets = [
      {
        position = {
          col   = 1
          row   = 1
          sizex = 6
          sizey = 4
        }
        config = {
          name        = "How to use dashboards"
          description = "HTML widget from template"
          type        = "html"
          theme       = "borderPurple"
          interval    = 5
          timescale   = "day"
          isCustom    = false
          resources = [
            {
              type = "html"
              URL  = "https://www.youtube.com/embed/-WNU4ffumk0"
            }
          ]
          version = 2
        }
      }
    ]
  })
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