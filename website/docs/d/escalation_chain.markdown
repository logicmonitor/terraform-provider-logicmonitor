---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_escalation_chain"
sidebar_current: "docs-logicmonitor-datasources-escalation-chain"
description: |-
  Get information on a LogicMonitor escalation chain resource
---

# logicmonitor_escalation_chain

This can be used to get information on a LogicMonitor escalation chain resource given a filter value from argument list

## Example Usage    
### EscalationChain
```hcl
# Datasource to get information of LogicMonitor escalation chain
data "logicmonitor_EscalationChain" "my_EscalationChain" {
        filter = "description~\"LM Escalation Chain testing\""
        depends_on = [
            logicmonitor_escalation_chain.myEscalationChain
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-escalation-chain. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

