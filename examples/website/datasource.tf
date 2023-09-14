resource "logicmonitor_website" "my_website"{
   type = "pingcheck"
   host = "google.com"
   overall_alert_level = "warn"
   polling_interval = 5
   description = "website test"
   disable_alerting = true
   stop_monitoring = true
   user_permission = "string"
   individual_sm_alert_enable = false
   steps = [
   {
    schema = "https"
    resp_type = "config"
    timeout = 0
    match_type = "plain"
    description = "string"
    use_default_root = true
    http_method = "GET"
    enable = true
    http_version = "1.1"
    follow_redirection = true
    post_data_edit_type = "raw"
    name = "string"
    req_type = "config"
    fullpage_load = false
   }
  ]
  transition = 1
  global_sm_alert_cond = 0
  is_internal = false
  domain = "www.ebay.com"
  name = "Ebay pingcheck"
  use_default_location_setting = false
  use_default_alert_setting = true
  individual_alert_level = "warn"
} 

data "logicmonitor_website" "my_website" {
   filter = "description~\"website test\""
 	depends_on = [
		logicmonitor_website.my_website
 	]
}

output "website" {
  description = "website"
  value       = data.logicmonitor_website.my_website.id
}
