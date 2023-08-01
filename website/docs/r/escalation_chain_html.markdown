---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_escalation_chain"
sidebar_current: "docs-logicmonitor-resource-escalation-chain"
description: |-
  Provides a LogicMonitor escalation chain resource. This can be used to create and manage LogicMonitor escalation chains.
---

# logicmonitor_escalation_chain

Provides a LogicMonitor escalation chain resource. This can be used to create and manage LogicMonitor escalation chains.

## Example Usage
```hcl
# Create a LogicMonitor escalation chain
resource "logicmonitor_escalation_chain" "my_escalation_chain" {
  description = "For alerts escalated to the NOC Team"
  destinations = [
    {
      period = [{
            week_days = 2
            timezone = "UTC"
            start_minutes = 10
            end_minutes   = 15
        }],
        stages = [{
            type    = "Admin"
            addr    = "unicornsparkles@rainbow.io"
            method  = "EMAIL"
            contact = "78362637"
        }]
        type = "timebased"
    }
  ]
  cc_destinations = [
    {
      method = "EMAIL"
      contact = "string"
      type = "Admin"
      addr = "unicornsparkles@rainbow.io"
    }
  ]
  enable_throttling = true
  name = "NOC Team"
  throttling_alerts = 40
  throttling_period = 30
}
```

## Argument Reference

The following arguments are **required**:
* `destinations` - 
  + `period` -  
    + `weekDays` - the list of week day of this period (required)
    + `timezone` - the timezone for this period (required)
    + `startMinutes` - the start minute of this period (required)
    + `endMinutes` - the end minute of this period (required)
  + `stages` - (required) 
    + `addr` - the user name if method = admin, or the email address if method = arbitrary
    + `contact` - contact details, email address or phone number
    + `method` (required) - Admin | Arbitrary, where Admin = a user, and Arbitrary = an arbitrary email
    + `type` (required) - email | sms | voice, where type must be email if method = arbitrary
  + `type` - single (required)
* `name` - the chain name

The following arguments are **optional**:
* `cc_destinations` - 
  + `addr` - the user name if method = admin, or the email address if method = arbitrary
  + `contact` - contact details, email address or phone number
  + `method` (required) - Admin | Arbitrary, where Admin = a user, and Arbitrary = an arbitrary email
  + `type` (required) - email | sms | voice, where type must be email if method = arbitrary
* `description` - 
* `enable_throttling` - if throttle needs to be enabled then true if not then false.
* `throttling_alerts` - max number of alert can send during a throttle period
* `throttling_period` - the throttle period

## Import

escalation chains can be imported using their escalation chain ID or name
```
$ terraform import logicmonitor_escalation_chain.my_escalation_chain 66
$ terraform import logicmonitor_escalation_chain.my_escalation_chain NOC Team
```