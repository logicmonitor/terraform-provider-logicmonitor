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
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/alert-rules/about-the-alert-rules-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

