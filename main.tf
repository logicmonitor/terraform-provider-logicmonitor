provider "logicmonitor" {
  api_id = "Q9pUaz94K29NanC2Rb97"
  api_key = "h7ItF67232Xw6-dQ9}6+T4vvh}9=dt9RqVM64G+!"
  company = "perrysportal"
}

resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.19"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
  hostgroup_id = "${data.logicmonitor_device_group.devicegroups.id}"
  properties {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}

resource "logicmonitor_device_group" "group1" {
    name = "NewGroupPerry"
    properties {
     "jmx.port" = "9003"
     "system.categories" = "ec2"
    }
}

resource "logicmonitor_collector_group" "group1" {
    name = "NewGroupPerry"
    properties {
     "jmx.port" = "9003"
     "system.categories" = "ec2"
    }
}

resource "logicmonitor_collector" "collector1" {
    description                       = "test collector"
    enable_collector_device_failover  = false
    enable_failback                   = false
    resend_interval                   = 5
    suppress_alert_clear              = false
    properties {
     "test" = "a"
    }
}

data "logicmonitor_collectors" "collectors" {
  filters {
    "custom_property_name" = "pod"
    "operator" = ":"
    "custom_property_value" = "p01_ld5_prod"
    }
}

data "logicmonitor_device_group" "devicegroups" {
  filters {
    "property" = "fullPath"
    "operator" = ":"
    "value" = "NewTestGroup/NewGroupTesting"
  }
}
