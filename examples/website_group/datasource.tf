resource "logicmonitor_website_group" "my_website_group"{
  stop_monitoring = false
  description = "website group test001"
  disable_alerting = false
  parent_id = 1
  name = "Amazon Website Checks"
  properties = [
  {
  name : "addr",
  value : "127.0.0.1"
  }  
  ]
}

data "logicmonitor_website_group" "my_website_group" {
   filter = "description~\"website group test\""
 	depends_on = [
		logicmonitor_website_group.my_website_group
 	]
}

output "websiteGroup" {
  description = "website group"
  value       = data.logicmonitor_website_group.my_website_group.id
}
