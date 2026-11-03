---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_alert_rule"
sidebar_current: "docs-logicmonitor-datasources-alert-rule"
description: |-
  Get information on a LogicMonitor alert rule resource
---

# logicmonitor_alert_rule

This can be used to get information on a LogicMonitor alert rule resource given a filter value from argument list

## Example Usage    
### AlertRule
```hcl
# Datasource to get information of LogicMonitor alert rule
data "logicmonitor_AlertRule" "my_AlertRule" {
        filter = "name~\"Alert Rule Testing\""
        depends_on = [
            logicmonitor_alert_rule.myAlertRule
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-alert-rules. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_alert_rules` returns every alert rule that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_alert_rules" "all_matches" {
  filter    = "name~\"Alert Rule Testing\""
  page_size = 300
}

output "alert_rule_count" {
  value = data.logicmonitor_alert_rules.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_alert_rule` above. Omit it to retrieve every alert rule.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of alert rule records matching `filter`.
* `items` - List of matching alert rule records. Each entry exposes the same read-only attributes as `logicmonitor_alert_rule` above, so e.g. `data.logicmonitor_alert_rules.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_alert_rule" "imported" {
  for_each = {
    for item in data.logicmonitor_alert_rules.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```