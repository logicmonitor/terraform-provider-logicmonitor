---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_website_group"
sidebar_current: "docs-logicmonitor-datasources-website-group"
description: |-
  Get information on a LogicMonitor website group resource
---

# logicmonitor_website_group

This can be used to get information on a LogicMonitor website group resource given a filter value from argument list

## Example Usage    
### WebsiteGroup
```hcl
# Datasource to get information of LogicMonitor website group
data "logicmonitor_WebsiteGroup" "my_WebsiteGroup" {
        filter = "description~\"website group test\""
 	    depends_on = [
		logicmonitor_website_group.my_website_group
 	    ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/service-groups/about-the-service-group-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

