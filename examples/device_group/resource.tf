resource "logicmonitor_device_group" "my_device_group" {
  description          = "normal device group test"
  disable_alerting     = true
  enable_netflow       = false
  default_collector_id = 10
  name                 = "meng test device group test"
  parent_id            = 1
  custom_properties {
    name  = "addr"
    value = "127.0.0.1"
  }
  group_type = "Normal"
  default_auto_balanced_collector_group_id = 0
}

/* AWS Device Group Testing */

/*
provider "aws" {
  access_key = ""
  secret_key = ""
  region     = "us-east-1"
}

data "logicmonitor_data_resource_aws_external_id" "my_external_id" {
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = [
      "sts:AssumeRole"
    ]
    condition {
      test = "StringEquals"
      values = [
        data.logicmonitor_data_resource_aws_external_id.my_external_id.external_id
      ]
      variable = "sts:ExternalId"
    }
    effect = "Allow"
    principals {
      identifiers = [
        "282028653949"
      ]
      type = "AWS"
    }
  }
}

resource "aws_iam_role" "lm" {
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
  name               = "TF-Integration-Role-Swag"
}

resource "aws_iam_role_policy_attachment" "read_only_access" {
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
  role       = aws_iam_role.lm.name
}

resource "aws_iam_role_policy_attachment" "aws_support_access" {
  policy_arn = "arn:aws:iam::aws:policy/AWSSupportAccess"
  role       = aws_iam_role.lm.name
}

// this resource acts as a timing buffer between the update of the AWS resources
// and the creation of the AWS Device Group in LM. This is required because the
// AWS API takes some time before the changes to the role is reflected. If done
// to quickly, the LM operation will fail.
resource "null_resource" "wait" {
  triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    command = "sleep 10"
  }
  // runs after all dependent AWS operations are complete
  depends_on = [
    aws_iam_role.lm,
    aws_iam_role_policy_attachment.read_only_access,
    aws_iam_role_policy_attachment.aws_support_access,
  ]
}

#Required fields: assumed_role_arn, external_id
resource "logicmonitor_device_group" "my_aws_device_group" {
  description      = "test description"
  disable_alerting = true
  enable_netflow   = false
  extra {
    account {
      assumed_role_arn = aws_iam_role.lm.arn
      external_id      = data.logicmonitor_data_resource_aws_external_id.my_external_id.external_id
    }
    default {
      disable_terminated_host_alerting = true
      select_all                       = false
      monitoring_region_infos          = ["US_EAST_1", "US_EAST_2", "US_WEST_1", "US_WEST_2"]
      dead_operation                   = "MANUALLY"
      use_default                      = true
      name_filter                      = []
      tags = [
        {
          operation = "include"
          name      = "system.aws.tag.test"
          value     = "localdev"
        },
        {
          operation = "exclude"
          name      = "system.aws.tag.test1"
          value     = "HighPriority"
        }
      ]
    }
    services {
      s_q_s {
        // because use_default is true, no other fields are needed
        use_default = true
      }
    }
  }
  group_type = "AWS/AwsRoot"
  name       = "Test AWS Group"
  parent_id  = 1
  // Must guarantee that the AWS resources are updated before creating the device group
  depends_on = [
    null_resource.wait,
  ]
}
*/

/*
output "my_external_id" {
	description = "My external ID"
	value =  data.logicmonitor_data_resource_aws_external_id.my_external_id.external_id
}

output "my_arn" {
	description = "My ARN"
	value =  aws_iam_role.lm.arn
}
*/


/* AZURE Device Group Testing */

/*
variable "user_info" {
  type = object({
    client_id  = string
    secret_key = string
    tenant_id  = string
  })
  nullable  = false
  sensitive = true
}

// Required fields: client_id, secret_key, tenant_id and subscription_ids
resource "logicmonitor_device_group" "my_azure_device_group" {
  name             = "Azure test"
  parent_id        = 1
  group_type       = "Azure/AzureRoot"
  description      = "Azure test"
  disable_alerting = false
  # enable_netflow = false
  custom_properties = [
    {
      name  = "customer"
      value = "customerA"
    }
  ]
  extra {
    account {
      schedule              = "0 * * * *"
      country               = "USA"
      client_id             = var.user_info.client_id
      secret_key            = var.user_info.secret_key
      subscription_ids      = ""
      tenant_id             = var.user_info.tenant_id
      collector_description = "Cloud Collector"
      # collector_id = -2
    }
    default {
      disable_terminated_host_alerting = true
      //normal_collector_config {}
      custom_n_s_p_schedule = ""
      select_all            = true
      monitoring_regions = [
        "CENTRAL_US",
        "EAST_US",
        "EAST_US_2"
      ]
      device_display_name_template        = ""
      disable_stop_terminate_host_monitor = true
      dead_operation                      = "KEEP_7_DAYS"
      use_default                         = true
      name_filter                         = []
      tags                                = []
    }
    services {
      v_i_r_t_u_a_l_m_a_c_h_i_n_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      s_t_o_r_a_g_e_a_c_c_o_u_n_t {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []

      }
      e_v_e_n_t_h_u_b {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []

      }
      a_p_p_s_e_r_v_i_c_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []

      }
      a_p_p_s_e_r_v_i_c_e_p_l_a_n {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t_vm {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      p_u_b_l_i_c_ip {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      sql_d_a_t_a_b_a_s_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      t_a_b_l_e_s_t_o_r_a_g_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      b_l_o_b_s_t_o_r_a_g_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      n_e_t_w_o_r_k_i_n_t_e_r_f_a_c_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      f_i_l_e_s_t_o_r_a_g_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      q_u_e_u_e_s_t_o_r_a_g_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      l_o_g_i_c_a_p_p_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      a_p_p_l_i_c_a_t_i_o_n_i_n_s_i_g_h_t_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      api_m_a_n_a_g_e_m_e_n_t {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      a_p_p_l_i_c_a_t_i_o_n_g_a_t_e_w_a_y {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      a_u_t_o_m_a_t_i_o_n_a_c_c_o_u_n_t {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      b_a_c_k_u_p_p_r_o_t_e_c_t_e_d_i_t_e_m_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      c_o_s_m_o_s_d_b {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      d_a_t_a_f_a_c_t_o_r_y {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      e_x_p_r_e_s_s_r_o_u_t_e_c_i_r_c_ui_t {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      k_e_y_v_a_u_l_t {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      l_o_a_d_b_a_l_a_n_c_e_r_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      f_i_r_e_w_a_l_l {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      l_o_a_d_b_a_l_a_n_c_e_r_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      l_o_g_a_n_a_l_y_t_i_c_s_w_o_r_k_s_p_a_c_e_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      p_o_s_t_g_r_e_sql {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      m_y_sql {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      f_r_o_n_t_d_o_o_r_s {
        # disable_terminated_host_alerting = true
        # normal_collector_config {
        #   collectors = []
        #   enable    = false
        # }
        # custom_n_s_p_schedule = ""
        # select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        # device_display_name_template        = ""
        # disable_stop_terminate_host_monitor = true
        # dead_operation                      = "KEEP_7_DAYS"
        use_default = true
        # name_filter                         = []
        # tags                                = []
      }
      s_y_n_a_p_s_e_w_o_r_k_s_p_a_c_e_s {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      c_o_g_n_i_t_i_v_e_s_e_a_r_c_h {
        disable_terminated_host_alerting = true
        normal_collector_config {
          collectors = []
          enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "CENTRAL_US",
          "EAST_US_2",
          "WEST_US"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }

    }
  }
}

data "logicmonitor_device_group" "myDeviceGroup" {
  filter = "description~\"Azure test\""
  depends_on = [
    logicmonitor_device_group.my_azure_device_group
  ]
}

output "deviceGroups" {
  description = "Azure device groups"
  value       = data.logicmonitor_device_group.myDeviceGroup
}
*/

# GCP Device Group Testing
/*
variable "user_info_gcp" {
  type = object({
    project_id          = string
    service_account_key = any  # JSON object for service account key
  })
  nullable  = false
  sensitive = true
}
resource "logicmonitor_device_group" "my_gcp_device_group" {
  name             = "GCP Production"
  parent_id        = 1
  group_type       = "GCP/GcpRoot"
  description      = "GCP Production Environment"
  disable_alerting = false
  custom_properties = [
   {
			name = "location"
      value = "pune"
	 },
   {
			name = "host"
      value = "localhost"
	 }
  ]
  extra {
    account {
      # GCP-specific account fields
      project_id              = var.user_info_gcp.project_id
      service_account_key     = jsonencode(var.user_info_gcp.service_account_key)
      schedule                = "0 * * * *"  # Hourly discovery
      collector_description   = "GCP Cloud Collector"
    }
    default {
      # Default settings for all GCP services
      use_default                         = true
      select_all                          = false
      monitoring_regions                  = ["us-central1", "us-west1", "europe-west1"]
      dead_operation                      = "KEEP_7_DAYS"
      disable_terminated_host_alerting    = true
      disable_stop_terminate_host_monitor = true
      name_filter = ["prod-*", "*-production"]
      tags = [
        {
          operation = "include"
          name      = "system.gcp.label.environment" 
          value     = "production"
        },
        {
          operation = "exclude"
          name      = "system.gcp.label.temporary"
          value     = "true"
        }
      ]
    }
    services {
      a_p_p_e_n_g_i_n_e {
        use_default     = true
      }
      c_l_o_u_d_a_ip_l_a_t_f_o_r_m {
        use_default     = true
      }
      c_l_o_u_d_b_i_g_t_a_b_l_e {
        use_default     = true
      }
      c_l_o_u_d_c_o_m_p_o_s_e_r {
        use_default     = true
      }
      c_l_o_u_d_d_a_t_a_f_l_o_w {
        use_default     = true
      }
      c_l_o_u_d_d_a_t_a_p_r_o_c {
        use_default     = true
      }
      c_l_o_u_d_d_l_p {
        use_default     = true
      }
      c_l_o_u_d_dns {
        use_default     = true
      }
      c_l_o_u_d_f_i_l_e_s_t_o_r_e {
        use_default     = true
      }
      c_l_o_u_d_f_i_r_e_s_t_o_r_e {
        use_default     = true
      }
      c_l_o_u_d_f_u_n_c_t_i_o_n {
        use_default     = true
      }
      c_l_o_u_d_i_n_t_e_r_c_o_n_n_e_c_t {
        use_default     = true
      }
      c_l_o_u_d_i_o_t {
        use_default     = true
      }
      c_l_o_udp_u_b_s_u_b {
        use_default     = true
      }
      c_l_o_u_d_r_e_d_i_s {
        use_default     = true
      }
      c_l_o_u_d_r_o_u_t_e_r {
        use_default     = true
      }
      c_l_o_u_d_r_u_n {
        use_default     = true
      }
      c_l_o_u_d_s_p_a_n_n_e_r {
        use_default     = true
      }
      c_l_o_u_d_sql {
        use_default     = true
      }
      c_l_o_u_d_s_t_o_r_a_g_e {
        use_default     = true
      }
      c_l_o_u_d_t_a_s_k_s {
        use_default     = true
      }
      c_l_o_u_d_t_p_u {
        use_default     = true
      }
      c_l_o_u_d_t_r_a_c_e {
        use_default     = true
      }
      c_o_m_p_u_t_e_e_n_g_i_n_e {
        use_default     = true
      }
      c_o_m_p_u_t_e_e_n_g_i_n_e_a_u_t_o_s_c_a_l_e_r {
        use_default     = true
      }
      https_l_o_a_d_b_a_l_a_n_c_e_r {
        use_default     = true
      }
      i_n_t_e_r_c_o_n_n_e_c_t_a_t_t_a_c_h_m_e_n_t {
        use_default     = true
      }
      m_a_n_a_g_e_d_s_e_r_v_i_c_e_f_o_r_m_i_c_r_o_s_o_f_t_a_d {
        use_default     = true
      }
      n_e_t_w_o_r_k_l_o_a_d_b_a_l_a_n_c_e_r {
        use_default     = true
      }
      r_e_g_i_o_n_a_l_https_l_o_a_d_b_a_l_a_n_c_e_r {
        use_default     = true
      }
      r_e_g_i_o_n_a_l_n_e_t_w_o_r_k_l_o_a_d_b_a_l_a_n_c_e_r {
        use_default     = true
      }
      v_p_n_g_a_t_e_w_a_y {
        use_default     = true
      }
      v_p_n_h_a_g_a_t_e_w_a_y {
        use_default     = true
      }
   }
 }
}
*/

# OCI Device Group Testing 
/*
variable "user_info_oci" {
  type = object({
    tenancy_id  = string
    user_id     = string
    private_key = string
  })
  nullable  = false
  sensitive = true
}
resource "logicmonitor_device_group" "my_OCI_device_group" {
  name             = "OCI Production"
  parent_id        = 1
  group_type       = "OCI/OciRoot"
  description      = "OCI Production Environment"
  disable_alerting = false
  custom_properties = [
   {
			name = "location"
      value = "pune"
	 },
   {
			name = "host"
      value = "localhost"
	 }
  ]
  extra {
    account {
      schedule              = "0 * * * *"
      country               = "USA"
      collector_description = "OCI Cloud Collector"
      tenancy_id             = var.user_info_oci.tenancy_id
      user_id                = var.user_info_oci.user_id
      private_key            = var.user_info_oci.private_key
    }
  default {
      # Default settings for all OCI services
      use_default                         = true
      select_all                          = false
      monitoring_regions                  = []
      dead_operation                      = "KEEP_7_DAYS"
      disable_terminated_host_alerting    = true
      disable_stop_terminate_host_monitor = true
    }
    services {
      o_c_i_b_l_o_c_k_s_t_o_r_a_g_e {
        disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_i_a_u_t_o_n_o_m_o_u_s_d_a_t_a_b_a_s_e{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_i_c_o_m_p_u_t_e{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_id_y_n_a_m_i_c_r_o_u_t_i_n_g_g_a_t_e_w_a_y{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_i_f_u_n_c_t_i_o_n_s{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_i_o_b_j_e_c_t_s_t_o_r_a_g_e{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
      o_c_i_s_i_t_e_t_o_s_i_t_e_v_p_n{
       disable_terminated_host_alerting = true
        normal_collector_config {
        collectors = []
        enable    = false
        }
        custom_n_s_p_schedule = ""
        select_all            = true
        monitoring_regions = [
          "us-phoenix-1"
        ]
        device_display_name_template        = ""
        disable_stop_terminate_host_monitor = true
        dead_operation                      = "KEEP_7_DAYS"
        use_default                         = true
        name_filter                         = []
        tags                                = []
      }
    }
  }
}
*/
