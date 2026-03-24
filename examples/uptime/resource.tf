variable "require_auth" {
  description = "Whether authentication is required"
  type        = bool
  default     = false
}
resource "logicmonitor_device" "my_device"{
	current_collector_id = 2
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
  display_name = "uptime webcheck"
  host_group_ids = 276147
  need_stc_grp_and_sorted_c_p = false
  preferred_collector_id = 2
  type = "uptimewebcheck"
  model = "websiteDevice"
  device_type = 18
  name = "Test-Web-Check-Int-V3-Int-New-111"
	description = ""
  group_ids = [64]
  resource_ids = [
		{
			name = "addr"
      value = "127.0.0.1"
		}
  ]
  is_internal = true
  global_sm_alert_cond = 0
  use_default_alert_setting = false
  use_default_location_setting = false
  test_location = [
    {
      all = true
      collector_ids = [2]
      smg_ids = []
    }
   ]
  page_load_alert_time_in_m_s = 30000
  alert_expr = ""
  domain = "www.google.com"
  host = "google.com"
  ignore_s_s_l = true
  schema = "https"
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