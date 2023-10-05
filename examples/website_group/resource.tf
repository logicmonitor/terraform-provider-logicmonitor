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