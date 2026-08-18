resource "logicmonitor_widget" "pie" {
  name         = "TS_PieChartWidget"
  description  = ""
  type         = "pieChart"
  dashboard_id = 65756
  theme        = "newBorderGray"
  interval     = 3
  timescale    = "day" # present in example TF config, not present in API payload
  pie_chart_info {
    title                       = "TS_PieChartWidget"
    show_labels_and_lines_on_pc = false
    max_slices_can_be_shown     = 10
    hide_zero_percent_slices    = false
    data_points {
      name                  = "dp0"
      device_group_full_path = "TS_DeviceGroup_dnd"
      device_display_name    = "TS_127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "MA_Test_Datasource_RT_KJ"
      data_point_name        = "dp0"
      aggregate              = false
      aggregate_function     = "SUM"
      type                   = "normal"
      top10                  = false
      glob_mode              = true
    }
    # counters / virtual_data_points are nested blocks — omit when empty
    # (do not use counters = [] or virtual_data_points = []).
    pie_chart_items {
      data_point_name = "dp0"
      legend          = "dp0"
      color           = "auto"
    }
    group_remaining_as_others = false
  }
  support_key_word = "Pie Chart"
}



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
    data_points {
      name                   = "dp0"
      device_group_full_path = "TS_DeviceGroup_dnd"
      device_display_name    = "TS_127.0.0.1_collector_84408_dnd"
      instance_name          = "*"
      data_source_full_name  = "MA_Test_Datasource_RT_KJ"
      data_point_name        = "dp0"
      aggregate_function     = "SUM"
      type                   = "normal"
      data_point_id          = 223830
      data_source_id         = 139635667
    }
    # counters / virtual_data_points are nested blocks — omit when empty
    big_number_items {
      data_point_name               = "dp0"
      bottom_label                  = "dp0"
      right_label                   = "dp0"
      position                      = 2
      rounding                      = 2
      use_comma_separators          = true
      change_threshold_color_toggle = false
    }
  }
}

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
    # virtual_data_points is a nested block — omit when empty
    data_points {
      name                  = "dp0"
      data_point_name       = "dp0"
      aggregate_function    = "SUM"
      data_source_full_name = "MA_Test_Datasource_RT_KJ"
      data_source_id        = 139635667
      consolidate_function  = "average"
      display = {
        option = "default"
      }
      device_group_full_path = {
        value   = "TS_DeviceGroup_dnd"
        is_glob = "false"
      }
      device_display_name = {
        value   = "TS_127.0.0.1_collector_84408_dnd"
        is_glob = "false"
      }
      instance_name = {
        value   = "*"
        is_glob = "true"
      }
      is_virtual_datapoint = true
      is_valid_exp         = true
    }
  }
}

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
    # alternate_data_points is a nested block — omit when empty
      alternate_data_points {
        data_source_full_name = ""
        data_source_id        = "124192717"
        data_point_name       = ""
        data_point_id         = 0
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
    clone_widget_id         = 334875
    widget_filters          = jsonencode({})
  }
}

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
    clone_widget_id         = 318171
    widget_filters          = jsonencode({})
  }
}

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
        },
        {
          columnKey   = "alert-device"
          columnLabel = "Resource/Website"
          visible     = true
          passthroughProps = {
            hasDisplayName = true
            hasIcon        = false
          }
          columnSize    = 389
          isSortable    = true
          componentName = "ForeignKeyCell"
          dataPath = [
            "monitoredObjectId"
          ]
        },
        {
          columnKey   = "alert-datasource"
          columnLabel = "LogicModule"
          visible     = true
          passthroughProps = {
            hasDisplayName = true
            hasIcon        = false
          }
          columnSize    = 138
          isSortable    = true
          componentName = "ForeignKeyCell"
          dataPath = [
            "moduleId"
          ]
        },
        {
          columnKey   = "alert-datasource-instance"
          columnLabel = "Instance"
          visible     = true
          passthroughProps = {
            hasDisplayName = true
            hasIcon        = false
          }
          columnSize    = 147
          isSortable    = true
          componentName = "ForeignKeyCell"
          dataPath = [
            "instanceId"
          ]
        },
        {
          columnKey     = "alert-datapoint"
          columnLabel   = "Datapoint"
          visible       = true
          columnSize    = 216
          isSortable    = true
          componentName = "SpaceId"
          dataPath = [
            "datapointId"
          ]
        },
        {
          columnKey     = "alert-value"
          columnLabel   = "Alert Value"
          visible       = true
          columnSize    = 120
          isSortable    = false
          componentName = "AlertValue"
          dataPath = [
            "alertValue"
          ]
        },
        {
          columnKey     = "alert-static-threshold"
          columnLabel   = "Static Threshold"
          visible       = true
          columnSize    = 120
          isSortable    = false
          componentName = "ThresholdCell"
          dataPath = [
            "threshold"
          ]
        },
        {
          columnKey     = "alert-dynamic-threshold"
          columnLabel   = "Dynamic Threshold"
          visible       = true
          columnSize    = 120
          isSortable    = false
          componentName = "ThresholdCell"
          dataPath = [
            "threshold"
          ]
        },
        {
          columnKey   = "alert-group"
          columnLabel = "Group"
          visible     = false
          passthroughProps = {
            hasIcon = false
            type    = "dg"
          }
          columnSize    = 120
          isSortable    = false
          componentName = "AlertWidgetGroupsCell"
          dataPath = [
            "monitoredObjectGroups"
          ]
        },
        {
          columnKey     = "alert-notes"
          columnLabel   = "Notes"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "notes"
          ]
        },
        {
          columnKey     = "alert-acked-by"
          columnLabel   = "ACK By"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "ACKBy"
          ]
        },
        {
          columnKey   = "alert-acked-on"
          columnLabel = "ACK At"
          visible     = false
          passthroughProps = {
            hasRelativeTime = false
          }
          columnSize    = 120
          isSortable    = false
          componentName = "SpaceDateTime"
          dataPath = [
            "ACKAtMS"
          ]
        },
        {
          columnKey   = "alert-cleared-on"
          columnLabel = "Cleared At"
          visible     = false
          passthroughProps = {
            hasRelativeTime = true
          }
          columnSize    = 120
          isSortable    = false
          componentName = "SpaceDateTime"
          dataPath = [
            "clearedAtMS"
          ]
        },
        {
          columnKey     = "alert-in-sdt"
          columnLabel   = "Is In SDT"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "isSDT"
          ]
        },
        {
          columnKey     = "alert-rule-name"
          columnLabel   = "Alert Rule"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "SpaceId"
          dataPath = [
            "ruleId"
          ]
        },
        {
          columnKey     = "alert-escalation-chain"
          columnLabel   = "Escalation Chain"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "SpaceId"
          dataPath = [
            "escalationChainId"
          ]
        },
        {
          columnKey     = "alert-datasource-instance-description"
          columnLabel   = "Instance Description"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "instanceDescription"
          ]
        },
        {
          columnKey     = "alert-full-path"
          columnLabel   = "Full Path"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "fullPath"
          ]
        },
        {
          columnKey   = "alert-routing-state"
          columnLabel = "Notification State"
          visible     = false
          columnSize  = 120
          isSortable  = false
          dataPath = [
            "dependencyRoutingState"
          ]
        },
        {
          columnKey     = "alert-dependency-role"
          columnLabel   = "Dependency Role"
          visible       = false
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "dependencyRole"
          ]
        },
        {
          columnKey     = "alert-dependent-alerts"
          visible       = false
          columnLabel   = "Dependent Alerts Count"
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "dependentAlertsCount"
          ]
        },
        {
          columnKey     = "alert-grouping-criteria"
          visible       = false
          columnLabel   = "Alert Grouping Criteria"
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "alertGroupEntityValue"
          ]
        },
        {
          columnKey     = "alert-log-metadata"
          visible       = false
          columnLabel   = "Log Metadata"
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "logMetaData"
          ]
        },
        {
          columnKey     = "alertExternalTicketUrl"
          columnLabel   = "ServiceNow incident"
          visible       = true
          columnSize    = 120
          isSortable    = false
          componentName = "SNIncidentCell"
          dataPath = [
            "alertExternalTicketUrl"
          ]
        },
        {
          columnKey     = "alert-log-partition"
          visible       = false
          columnLabel   = "Logs Partition"
          columnSize    = 120
          isSortable    = false
          componentName = "Body1"
          dataPath = [
            "logPartition"
          ]
        },
        {
          columnKey     = "alert-historical-sdt"
          columnLabel   = "Historical SDT"
          visible       = true
          columnSize    = 120
          isSortable    = false
          componentName = "HistoricalSdtCell"
          dataPath = [
            "sdtIds"
          ]
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
    clone_widget_id         = 318167
  }
  support_key_word = "Alert List"
}
