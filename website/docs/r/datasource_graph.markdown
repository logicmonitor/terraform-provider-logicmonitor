---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_datasource_graph"
sidebar_current: "docs-logicmonitor-resource-datasource-graph"
description: |-
  Manages the graphs and overview graphs of a LogicMonitor datasource.
---

# logicmonitor_datasource_graph / logicmonitor_datasource_overview_graph

Graphs (`/setting/datasources/{id}/graphs`) and overview graphs (`/setting/datasources/{id}/ographs`) of a datasource.
Both resources take the same arguments; overview graphs also have `aggregated` and a per-datapoint `aggregate_method`.
Graph datapoints name the datasource datapoint they show; the provider looks up its id.

## Example Usage
```hcl
resource "logicmonitor_datasource_graph" "cpu_time" {
  datasource_id  = logicmonitor_datasource.cpu.id
  name           = "CPU Time"
  title          = "CPU Time"
  vertical_label = "seconds"
  data_points {
    name = "cpuIdleSeconds"
  }
  data_points {
    name = "cpuUserSeconds"
  }
  virtual_data_points {
    name = "cpuBusySeconds"
    rpn  = "cpuUserSeconds+cpuSystemSeconds"
  }
  lines {
    data_point_name = "cpuIdleSeconds"
    color_name      = "gray"
  }
  lines {
    data_point_name       = "cpuBusySeconds"
    is_virtual_data_point = true
  }
}

resource "logicmonitor_datasource_overview_graph" "cpu_busy" {
  datasource_id  = logicmonitor_datasource.cpu.id
  name           = "CPU Busy Percent"
  title          = "CPU Busy Percent"
  vertical_label = "%"
  data_points {
    name             = "cpuBusyPercent"
    aggregate_method = "sum"
  }
  lines {
    data_point_name = "cpuBusyPercent"
    legend          = "##INSTANCE##"
  }
}
```

## Argument Reference

The following arguments are **required**:
* `datasource_id` - The datasource the graph belongs to (string)
* `name` - The graph name (string)
* `title` - The graph title (string)
* `vertical_label` - The y-axis label (string)

The following arguments are **optional** (left out, the value LogicMonitor sets is kept):
* `time_scale` (e.g. 1day), `width`, `height`, `display_prio`, `rigid`, `base1024`
* `min_value`, `max_value` - numbers, or "NaN" for automatic (string)
* `aggregated` - overview graphs only (bool)
* `data_points` - set of `{name, data_point_name, consolidate_func, aggregate_method}`. `name` is the datasource datapoint;
  `data_point_name` is only needed when the graph datapoint has a different name. `consolidate_func` defaults to 1 (average);
  `aggregate_method` (e.g. sum, average) is for overview graphs.
* `virtual_data_points` - set of `{name, rpn}`
* `lines` - set of `{data_point_name, legend, type, color_name, is_virtual_data_point}`; `type` defaults to 1, `color_name` to "Auto".
  LogicMonitor does not keep the line order.

## Import

Graphs are imported as `<datasource id>/<graph id>`
```
$ terraform import logicmonitor_datasource_graph.cpu_time 24873941/27236
$ terraform import logicmonitor_datasource_overview_graph.cpu_busy 24873941/8844
```
