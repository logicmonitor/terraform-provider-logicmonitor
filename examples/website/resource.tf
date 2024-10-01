variable "require_auth" {
  description = "Whether authentication is required"
  type        = bool
  default     = true
}
resource "logicmonitor_website" "my_website"{
   type = "webcheck"
   host = "google.com"
   overall_alert_level = "warn"
   polling_interval = 5
   description = "website test"
   disable_alerting = true
   stop_monitoring = true
   schema = "http"
   user_permission = "string"
   test_location = [
    {
      all = false
      collector_ids = [1, 2, 3]
      collectors = null
      smg_ids = [85, 90]
    }
   ]
   group_id = 8
   individual_sm_alert_enable = false
   steps = [
   {
    schema = "https"
    resp_type = "config"
    timeout = 0
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
  transition = 1
  global_sm_alert_cond = 0
  is_internal = false
  domain = "www.example.com"
  name = "Ebay webcheck test"
  use_default_location_setting = false
  use_default_alert_setting = true
  individual_alert_level = "warn"
  trigger_s_s_l_expiration_alert = true
  trigger_s_s_l_status_alert = true
  ignore_s_s_l = false
  alert_expr = "< 100"
} 