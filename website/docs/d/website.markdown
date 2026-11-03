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


## Retrieving multiple matches

`logicmonitor_websites` returns every website that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_websites" "all_matches" {
  filter    = "description~\"website test\""
  page_size = 300
}

output "website_count" {
  value = data.logicmonitor_websites.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_website` above. Omit it to retrieve every website.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of website records matching `filter`.
* `items` - List of matching website records. Each entry exposes the same read-only attributes as `logicmonitor_website` above, so e.g. `data.logicmonitor_websites.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_website" "imported" {
  for_each = {
    for item in data.logicmonitor_websites.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```