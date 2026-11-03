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


## Retrieving multiple matches

`logicmonitor_escalation_chains` returns every escalation chain that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_escalation_chains" "all_matches" {
  filter    = "description~\"LM Escalation Chain testing\""
  page_size = 300
}

output "escalation_chain_count" {
  value = data.logicmonitor_escalation_chains.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_escalation_chain` above. Omit it to retrieve every escalation chain.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of escalation chain records matching `filter`.
* `items` - List of matching escalation chain records. Each entry exposes the same read-only attributes as `logicmonitor_escalation_chain` above, so e.g. `data.logicmonitor_escalation_chains.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_escalation_chain" "imported" {
  for_each = {
    for item in data.logicmonitor_escalation_chains.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```