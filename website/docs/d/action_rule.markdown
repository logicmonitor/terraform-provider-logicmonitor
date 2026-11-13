---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_action_rule"
sidebar_current: "docs-logicmonitor-datasources-action-rule"
description: |-
  Get information on a LogicMonitor action rule resource
---

# logicmonitor_action_rule

This can be used to get information on a LogicMonitor action rule resource given a filter value from argument list

## Example Usage    
### ActionRule
```hcl
# Datasource to get information of LogicMonitor action rule
data "logicmonitor_ActionRule" "my_ActionRule" {
        filter = "name~\"action Rule1\""
        depends_on = [
            logicmonitor_action_rule.my_action_rule
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_action_rules` returns every action rule that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_action_rules" "all_matches" {
  filter    = "name~\"action Rule1\""
  page_size = 300
}

output "action_rule_count" {
  value = data.logicmonitor_action_rules.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_action_rule` above. Omit it to retrieve every action rule.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of action rule records matching `filter`.
* `items` - List of matching action rule records. Each entry exposes the same read-only attributes as `logicmonitor_action_rule` above, so e.g. `data.logicmonitor_action_rules.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_action_rule" "imported" {
  for_each = {
    for item in data.logicmonitor_action_rules.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```