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


data "logicmonitor_widget" "big_number" {
   filter = "description~\"This is a test\""
 	depends_on = [
		logicmonitor_widget.big_number
 	]
}

output "widget" {
  description = "widget"
  value       = data.logicmonitor_widget.big_number.id
}
