resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
  hostgroup_id = "${logicmonitor_device_group.group1.id}"
  properties {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}

resource "logicmonitor_device_group" "test" {
  name = "Test"
}

resource "logicmonitor_device" "host1" {
  ip_addr = "mesos02.us-east-1.logicmonitor.net"
  display_name = "mesos02.us-east-1"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
}

resource "logicmonitor_device_group" "group1" {
    name = "NewGroup"
    properties {
     "jmx.port" = "9003"
     "system.categories" = "ec2"
    }
}

data "logicmonitor_collectors" "collectors" {
  most_recent = true
}
