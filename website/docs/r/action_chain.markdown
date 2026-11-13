---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_action_chain"
sidebar_current: "docs-logicmonitor-resource-action-chain"
description: |-
  Provides a LogicMonitor action chain resource. This can be used to create and manage LogicMonitor action chains.
---

# logicmonitor_action_chain
Provides a LogicMonitor action chain resource. This can be used to create and manage LogicMonitor action chains.

## Example Usage
```hcl
# Create a LogicMonitor action chain
resource "logicmonitor_action_chain" "my_action_chain" {
  description = "This chain description"
  name = "Test AC"
  stages {
    id   = 101
    type = "diagnosticSource"
  }
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The action chain name
   (string)
* `stages` - The list of stages in the action chain
   ([]*Stage)

The following arguments are **optional**:
* `description` - The action chain description (string)

## Import

action chains can be imported using their action chain ID or name
```
$ terraform import logicmonitor_action_chain.my_action_chain 66
$ terraform import logicmonitor_action_chain.my_action_chain Test AC
```