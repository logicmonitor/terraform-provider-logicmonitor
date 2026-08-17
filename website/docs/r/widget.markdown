---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_widget"
sidebar_current: "docs-logicmonitor-resource-widget"
description: |-
  Provides a LogicMonitor widget resource. This can be used to create and manage LogicMonitor widgets.
---

# logicmonitor_widget
Provides a LogicMonitor widget resource. This can be used to create and manage LogicMonitor widgets.

## Example Usage

The `logicmonitor_widget` resource is polymorphic. Set `type` to `cgraph`, `bigNumber`, `pieChart`, `table`, `noc`, or `alert` and provide the matching configuration block (`graph_info`, `big_number_info`, `pie_chart_info`, `table_info`, `noc_info`, or `alert_info`).

### Custom Graph Widget (`cgraph`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "custom_graph" {
  name             = "TS_CustomGraphWiget_Testing_update"
  description      = "This is a test"
  type             = "cgraph"
  dashboard_id     = 48860
  theme            = "newBorderGray"
  interval         = 10
  timescale        = "1day"
  support_key_word = "Custom Graph"

  graph_info {
    aggregate                   = true
    max_value                   = "NaN"
    min_value                   = "NaN"
    vertical_label              = "ms"
    top_x                       = 0
    desc                        = true
    global_consolidate_function = "VaST"
    scale_unit                  = 1000

    virtual_data_points {
      name = "vdp0"
      rpn  = "dp0 + dp1"
      display = {
        option = "default"
        type   = "line"
      }
    }

    data_points {
      name                  = "dp0"
      data_point_name       = "dp0"
      aggregate_function    = "SUM"
      data_source_full_name = "Test_Datasource"
      data_source_id        = 139635667
      consolidate_function  = "average"
      is_virtual_datapoint  = false
      is_valid_exp          = true

      display = {
        option = "default"
      }
      device_group_full_path = {
        value   = "DeviceGroup_dnd"
        is_glob = "false"
      }
      device_display_name = {
        value   = "127.0.0.1_collector_84408_dnd"
        is_glob = "false"
      }
      instance_name = {
        value   = "*"
        is_glob = "true"
      }
    }
  }
}
```

### Big Number Widget (`bigNumber`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "big_number" {
  name             = "TS_BigNumberWidget_updated"
  description      = "This is a test"
  type             = "bigNumber"
  dashboard_id     = 48860
  theme            = "newBorderGray"
  interval         = 10
  timescale        = "day"
  support_key_word = "Big Number"

  big_number_info {
    virtual_data_points {
      name = "vdp0"
      rpn  = "dp0 + dp1"
    }

    data_points {
      name                   = "dp0"
      device_group_full_path = "DeviceGroup_dnd"
      device_display_name    = "127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "Test_Datasource"
      data_point_name        = "dp0"
      aggregate_function     = "SUM"
      type                   = "normal"
      data_point_id          = 223830
      data_source_id         = 139635667
    }

    data_points {
      name                   = "dp1"
      device_group_full_path = "DeviceGroup_dnd"
      device_display_name    = "127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "Test_Datasource"
      data_point_name        = "dp1"
      aggregate_function     = "MAX"
      type                   = "normal"
      data_point_id          = 223831
      data_source_id         = 139635667
    }

    big_number_items {
      data_point_name               = "dp0"
      bottom_label                  = "dp0"
      right_label                   = "dp0"
      position                      = 1
      rounding                      = 2
      use_comma_separators          = true
      change_threshold_color_toggle = false
    }

    big_number_items {
      data_point_name               = "vdp0"
      bottom_label                  = "combined"
      right_label                   = ""
      position                      = 2
      rounding                      = 2
      use_comma_separators          = true
      change_threshold_color_toggle = false
    }
  }
}
```

### Pie Chart Widget (`pieChart`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "pie" {
  name             = "PieChartWidget"
  description      = ""
  type             = "pieChart"
  dashboard_id     = 65756
  theme            = "newBorderGray"
  interval         = 3
  timescale        = "day"
  support_key_word = "Pie Chart"

  pie_chart_info {
    title                       = "PieChartWidget"
    show_labels_and_lines_on_pc = false
    max_slices_can_be_shown     = 10
    hide_zero_percent_slices    = false
    group_remaining_as_others   = false

    virtual_data_points {
      name = "vdp0"
      rpn  = "dp0 + dp1"
    }

    data_points {
      name                   = "dp0"
      device_group_full_path = "DeviceGroup_dnd"
      device_display_name    = "127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "Test_Datasource"
      data_point_name        = "dp0"
      aggregate              = false
      aggregate_function     = "SUM"
      type                   = "normal"
      top10                  = false
      glob_mode              = true
    }

    data_points {
      name                   = "dp1"
      device_group_full_path = "DeviceGroup_dnd"
      device_display_name    = "127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "Test_Datasource"
      data_point_name        = "dp1"
      aggregate              = false
      aggregate_function     = "SUM"
      type                   = "normal"
      top10                  = false
      glob_mode              = true
    }

    pie_chart_items {
      data_point_name = "dp0"
      legend          = "dp0"
      color           = "auto"
    }

    pie_chart_items {
      data_point_name = "dp1"
      legend          = "dp1"
      color           = "auto"
    }

    pie_chart_items {
      data_point_name = "vdp0"
      legend          = "combined"
      color           = "auto"
    }
  }
}
```

### Table Widget (`table`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "table" {
  name            = "ctable"
  description     = ""
  type            = "table"
  dashboard_id    = 65756
  theme           = "newBorderGray"
  user_permission = "write"
  interval        = 3
  timescale       = "day"
  table_info {
    display_settings = jsonencode({
      propertyColumns = []
      pageSize        = "25"
      overrideFilter  = false
      showFilter      = false
      columnsV4 = [
        {
          columnName = "recvdpktsq"
          dataPoint = {
            dataSourceFullName = "Ping"
            dataSourceId       = "122104301"
            dataPointName      = "recvdpkts"
          }
          alternateDataPoints = [
            {
              dataSourceFullName = ""
              dataSourceId       = "124192717"
              dataPointName      = ""
            }
          ]
          unitLabel         = ""
          enableForecast    = false
          roundingDecimal   = 2
          rpn               = ""
          position          = 0
          propertiesOptions = "1"
          id                = "undefined_0"
          index             = 0
        }
      ]
    })
    is_support_custom_property = false
    columns {
      column_name = "recvdpktsq"
      data_point = {
        dataSourceFullName = "Ping"
        dataSourceId       = "122104301"
        dataPointName      = "recvdpkts"
      }
      alternate_data_points {
        data_source_full_name = ""
        data_source_id        = "124192717"
        data_point_name       = ""
      }
      unit_label         = ""
      enable_forecast    = false
      rounding_decimal   = 2
      rpn                = ""
      position           = 0
      properties_options = "1"
      id                 = "undefined_0"
      index              = 0
    }
    rows {
      label               = "device21"
      group_id            = 215809
      group_full_path     = ""
      device_id           = 1042326
      device_display_name = null
      instances {
        instance_id            = 0
        instance_name          = null
        data_point_name        = null
        data_point_id          = 0
        validation_status_code = 1
      }
      position = 0
      id       = "device21_0"
    }
    forecast = jsonencode({
      confidence = 70
      algorithm  = "ARIMA"
    })
    support_custom_property = false
    widget_filters          = jsonencode({})
  }
}
```

### NOC Widget (`noc`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "noc" {
  name            = "NOC"
  description     = ""
  type            = "noc"
  dashboard_id    = 65756
  theme           = "newBorderGray"
  user_permission = "write"
  interval        = 3
  timescale       = "day"
  noc_info {
    display_settings = jsonencode({
      showTypeIcon   = true
      displayAs      = "table"
      overrideFilter = false
      showFilter     = false
    })
    is_support_custom_property = false
    items {
      type                     = "device"
      device_group_full_path   = "gayatri-qapr-201"
      device_display_name      = "*"
      data_source_display_name = "Ping"
      instance_name            = "*"
      data_point_name          = "recvdpkts"
      group_by                 = "device"
      name                     = "##RESOURCENAME##"
      noc_item_options         = "1"
    }
    sort_by                 = "alertSeverity"
    display_column          = 1
    display_warn_alert      = true
    display_error_alert     = true
    display_critical_alert  = true
    ack_checked             = true
    sdt_checked             = true
    support_custom_property = false
    widget_filters          = jsonencode({})
  }
}
```

### Alert Widget (`alert`)
#### Omit fields entirely when empty (do not set an empty list).
```hcl
resource "logicmonitor_widget" "alert" {
  name            = "alerts1"
  description     = ""
  type            = "alert"
  dashboard_id    = 38044
  theme           = "newSolidDarkBlue"
  user_permission = "write"
  interval        = 3
  timescale       = "day"
  alert_info {
    display_settings = jsonencode({
      columns = [
        {
          columnKey   = "alert-began"
          columnLabel = "Reported At"
          visible     = true
          passthroughProps = {
            hasRelativeTime = true
          }
          columnSize    = 146
          isSortable    = true
          componentName = "SpaceDateTime"
          dataPath = [
            "reportedAtMS"
          ]
        },
        {
          columnKey     = "alert-severity"
          columnLabel   = "Severity"
          visible       = true
          columnSize    = 138
          isSortable    = true
          minSize       = 82
          componentName = "SeverityCell"
        }
      ]
      showFilter = false
      fontsize   = "normal-font"
      sort       = "-monitorObjectName"
      playSound = {
        criticalAlertAudioFileName = "none"
        errorAlertAudioFileName    = "none"
        warningAlertAudioFileName  = "none"
        shouldPlay                 = false
      }
      isShowAll = false
    })
    is_support_custom_property = false
    filters = {
      group                  = ""
      host                   = ""
      dataSource             = "Port-"
      instance               = "*"
      dataPoint              = ""
      severity               = ""
      acked                  = "all"
      sdted                  = "all"
      isHistoricalSdt        = "all"
      rule                   = ""
      chain                  = ""
      cleared                = "all"
      isEscalation           = "all"
      keyword                = ""
      dependencyRole         = ""
      dependencyRoutingState = ""
      anomaly                = ""
      suppressionType        = ""
    }
    parsed_filters = {
      group                  = ""
      host                   = ""
      dataSource             = "Port-"
      instance               = "*"
      dataPoint              = ""
      severity               = ""
      acked                  = "all"
      sdted                  = "all"
      rule                   = ""
      chain                  = ""
      isEscalation           = "all"
      isHistoricalSdt        = "all"
      cleared                = "all"
      keyword                = ""
      dependencyRole         = ""
      dependencyRoutingState = ""
      suppressionType        = ""
      anomaly                = ""
    }
    support_custom_property = false
    refresh_frequency       = jsonencode({ units = "SECONDS", offset = 180 })
    widget_color_schema     = "newSolidDarkBlue"
  }
  support_key_word = "Alert List"
}
```

## Argument Reference

The following arguments are **required**:

* `dashboard_id` - The id of the dashboard the widget belongs to (int32)
* `name` - The name of the widget (string)
* `type` - Widget type. Supported values: `cgraph` (custom graph), `bigNumber`, `pieChart` (string)

The following arguments are **optional**:

* `description` - The description of the widget (string)
* `interval` - The refresh interval of the widget, in minutes (int32)
* `theme` - The color scheme of the widget (string)
* `timescale` - The default timescale of the widget (string)
* `support_key_word` - Support keyword associated with the widget type (string)

Exactly one type-specific block must be set to match `type`:

* `graph_info` - Configuration for custom graph widgets (`type = "cgraph"`). Supports `data_points` (including `is_virtual_datapoint`, `is_valid_exp`, and TypeMap fields such as `device_display_name` / `instance_name` / `display`), `virtual_data_points`, `vertical_label`, `top_x`, `aggregate`, `min_value`, `max_value`, `scale_unit`, `global_consolidate_function`, and related graph settings.
* `big_number_info` - Configuration for big number widgets (`type = "bigNumber"`). Requires `data_points` and `big_number_items` blocks. Datapoints support `type`; items support `change_threshold_color_toggle`. Optional `counters` and `virtual_data_points` blocks may be omitted when empty.
* `pie_chart_info` - Configuration for pie chart widgets (`type = "pieChart"`). Requires `pie_chart_items`; `data_points` should be set for datapoint-backed slices and supports `type` / `glob_mode`. Optional `counters` and `virtual_data_points` blocks may be omitted when empty.

The following attributes are **exported**:

* `id` - The widget ID (int32)
* `last_updated_by` - The user that last updated the widget (string)
* `last_updated_on` - Last update time in epoch format (int64)
* `user_permission` - Permission level of the user who last modified the widget (string)

## Import

widgets can be imported using their widget ID or name
```
$ terraform import logicmonitor_widget.my_widget 66
$ terraform import logicmonitor_widget.my_widget Test
```