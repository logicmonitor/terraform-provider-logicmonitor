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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
          enabled    = false
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
