---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_action_chain"
sidebar_current: "docs-logicmonitor-datasources-action-chain"
description: |-
  Get information on a LogicMonitor action chain resource
---

# logicmonitor_action_chain

This can be used to get information on a LogicMonitor action chain resource given a filter value from argument list

## Example Usage    
### ActionChain
```hcl
# Datasource to get information of LogicMonitor action chain
data "logicmonitor_ActionChain" "my_ActionChain" {
        filter = "name~\"Test AC\""
        depends_on = [
            logicmonitor_action_chain.my_action_chain
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_action_chains` returns every action chain that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_action_chains" "all_matches" {
  filter    = "name~\"Test AC\""
  page_size = 300
}

output "action_chain_count" {
  value = data.logicmonitor_action_chains.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_action_chain` above. Omit it to retrieve every action chain.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of action chain records matching `filter`.
* `items` - List of matching action chain records. Each entry exposes the same read-only attributes as `logicmonitor_action_chain` above, so e.g. `data.logicmonitor_action_chains.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_action_chain" "imported" {
  for_each = {
    for item in data.logicmonitor_action_chains.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```