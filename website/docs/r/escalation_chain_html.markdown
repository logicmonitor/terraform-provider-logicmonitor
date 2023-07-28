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
  enable_throttling = true
  name = "NOC Team"
  throttling_alerts = 40
  throttling_period = 30
}
```

## Argument Reference

The following arguments are **required**:
* `destinations` - 
* `name` - 

The following arguments are **optional**:
* `cc_destinations` - 
* `description` - 
* `enable_throttling` - 
* `throttling_alerts` - 
* `throttling_period` - 

## Import

escalation chains can be imported using their escalation chain ID or name
```
$ terraform import logicmonitor_escalation_chain.my_escalation_chain 66
$ terraform import logicmonitor_escalation_chain.my_escalation_chain NOC Team
```