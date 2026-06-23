variable "require_auth" {
  description = "Whether authentication is required"
  type        = bool
  default     = false
}
resource "logicmonitor_device" "my_device"{
  properties = []
	custom_properties = [
		{
			name = "addr"
      value = "127.0.0.1"
		},
		{
			name = "host"
      value = "localhost"
		},
    {
			name = "location"
      value = "pune"
		},
    {
      name  = "system.categories"
      value = ""
    },
  ]
  type = "uptimewebcheck"
  device_type = 18
  name = "Test-Web-Check-Int-V3-Int-New-112201"
  display_name = "Test-Web-Check-Int-V3-Int-New-112201"
  preferred_collector_id = 10
	description = ""
  group_ids = [276147]
  is_internal = true
  global_sm_alert_cond = 0
  use_default_alert_setting = false
  use_default_location_setting = false
  overall_alert_level = "warn"
  polling_interval = 5
  test_location = [
    {
      all = true
      collector_ids = [30329]
      smg_ids = []
    }
   ]
  page_load_alert_time_in_m_s = 30000
  alert_expr = ""
  domain = "www.google.com"
  ignore_s_s_l = true
  schema = "https"
  transition = 1
  trigger_s_s_l_expiration_alert = false
  trigger_s_s_l_status_alert = false
  steps = [
   {
    schema = "https"
    resp_type = "config"
    timeout = 1
    match_type = "json"
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
    require_auth = var.require_auth
    http_headers = "X-Version:3"
    auth = var.require_auth ? [{
    password  = "string"
    type      = "basic"
    domain    = "string"
    user_name = "string"
    }] : []
    path = "$.data.resultKey"
    keyword = "DEVWRT-SANRT-JOB1-9127"
    http_body = "string"
    resp_script = "string"
    req_script = "string"
    label = "string"
    url = "/santaba/rest/version"
    type = "string"
    invert_match = true
    status_code = "200"
   }
  ]
}
data "logicmonitor_device" "my_devices" {
    filter = "displayName~\"uptime webcheck\""
	depends_on = [
		logicmonitor_device.my_device
	]
}

output "devices" {
  description = "devices"
  value       = data.logicmonitor_device.my_devices.id
}

