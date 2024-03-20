---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_website"
sidebar_current: "docs-logicmonitor-datasources-website"
description: |-
  Get information on a LogicMonitor website resource
---

# logicmonitor_website

This can be used to get information on a LogicMonitor website resource given a filter value from argument list

## Example Usage    
### Website
```hcl
# Datasource to get information of LogicMonitor website
data "logicmonitor_Website" "my_Website" {
        filter = "description~\"website test\""
 	    depends_on = [
		logicmonitor_website.my_website
 	    ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/services/about-the-service-api-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

