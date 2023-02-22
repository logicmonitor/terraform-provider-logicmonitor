---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device_group"
sidebar_current: "docs-logicmonitor-resource-device-group"
description: |-
  Provides a LogicMonitor device group resource. This can be used to create and manage LogicMonitor device groups.
---

# logicmonitor_device_group

Provides a LogicMonitor device group resource. This can be used to create and manage LogicMonitor device groups.

## Example Usage
```hcl
# Create a LogicMonitor device group
resource "logicmonitor_device_group" "my_device_group" {
  applies_to = "isLinux()"
  custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		}
	]
  default_collector_id = 1
  description = "Linux Servers"
  disable_alerting = true
  enable_netflow = true
  group_type = "Normal"
  name = "Linux Servers"
  parent_id = 1
}
```
## Example Usage - AWS Cloud Account
```hcl
provider "aws" {
  // add your AWS credentials
  access_key = ""
  secret_key = ""
  region     = ""
}

// used for verifying LM access to your AWS account
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
  name               = "TF-Integration-Role"
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
    command = "sleep 10s"
  }
  // runs after all dependent AWS operations are complete
  depends_on = [
	  aws_iam_role.lm,
	  aws_iam_role_policy_attachment.read_only_access,
	  aws_iam_role_policy_attachment.aws_support_access,
  ]
}

resource "logicmonitor_device_group" "my_aws_device_group" {
	description = "test description"
	disable_alerting = true
	enable_netflow = false
	extra {
		account {
			assumed_role_arn = aws_iam_role.lm.arn
			external_id = data.logicmonitor_data_resource_aws_external_id.my_external_id.external_id
		}
		default {
			disable_terminated_host_alerting = true
			select_all = false
			monitoring_region_infos = ["US_EAST_1", "US_EAST_2", "US_WEST_1", "US_WEST_2"]
			dead_operation = "MANUALLY"
			use_default = true
			name_filter = []
			tags = [
				{
					operation = "include"
					name = "system.aws.tag.test"
					value = "localdev"
				},
				{
					operation = "exclude"
					name = "system.aws.tag.test1"
					value = "HighPriority"
				}
			]
		}
		services {
			s_q_s {
        // because use_default is true, no other fields are needed
        // they settings will be determined by the above default object
				use_default = true
			}
		}
	}
	group_type = "AWS/AwsRoot"
	name = "Test AWS Group TF1"
	parent_id = 1
	// Must guarantee that the AWS resources are updated before creating the device group
	depends_on = [
		null_resource.wait,
	]
}
```
## Example Usage - Azure Cloud Account
```hcl
//Auth information for Azure account
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


```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the device group

The following arguments are **optional**:
* `applies_to` - The Applies to custom query for this group (only for dynamic groups)
* `custom_properties` - The properties associated with this device group
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `default_collector_id` - The Id of the default collector assigned to the device group
* `description` - The description of the device group
* `disable_alerting` - Indicates whether alerting is disabled (true) or enabled (false) for this device group
* `enable_netflow` - Indicates whether Netflow is enabled (true) or disabled (false) for the device group, the default value is true
* `extra` - The extra setting for cloud group
  + `account` - cloud account information (currently only supports AWS)
    + `accountId` - LogicMonitor's Account ID
    + `assumedRoleArn` - ARN of the role created for LogicMonitor to use while monitoring AWS (required)
    + `billingBucketName` - S3 bucket name that houses AWS billing info (deprecated) (computed)
    + `billingPathPrefix` - Path to billing info in billing bucket (deprecated) (computed)
    + `collectorId` - ID of the collector assigned to this group (computed)
    + `externalId` - External ID provide by LM for the creation of the assumed role in AWS (required for AWS)
    + `clientId`   - Client id received from Azure account (required for Azure)
    + `secretKey` - Secret key received from Azure Account (required for Azure)
    + `subscriptionIds` -  List of subcription ids(comma separated, required for Azure)) 
    + `tenantId` - Tenant id from Azure accountn (required for Azure)
    + `schedule` - The NetScan schedule for how frequently the cloud collector should scan/discover new resources in the cloud account. It's format is similar to Linux crontab but doesn't support some complex representations ('-', '/', ',') supported in standard linux crontabs.\nFormat: '*(minute) *(hour) *(day) *(week of month) *(weekday)'\nExamples: '50 * * * *' means scheduling at 50th minute per hour\n'50 10 20 * *' means scheduling at 10:50 of the 20th day per month\n'50 10 * 1 3' means scheduling at wednesday of the first week per month
    + `type` - Type of cloud account (computed)
  + `default` - default cloud service monitoring settings
    + `customNSPSchedule` - The NetScan schedule for how frequently the cloud collector should scan/discover new resources in the cloud account. It's format is similar to Linux crontab but doesn't support some complex representations ('-', '/', ',') supported in standard linux crontabs.\nFormat: '*(minute) *(hour) *(day) *(week of month) *(weekday)'\nExamples: '50 * * * *' means scheduling at 50th minute per hour\n'50 10 20 * *' means scheduling at 10:50 of the 20th day per month\n'50 10 * 1 3' means scheduling at wednesday of the first week per month
    + `deadOperation` - How to handle dead cloud devices (deletion from monitoring). Valid options include MANUALLY, IMMEDIATELY, KEEP_1_DAY, KEEP_3_DAYS, KEEP_7_DAYS, KEEP_14_DAYS, and KEEP_30_DAYS
    + `deviceDisplayNameTemplate` - Template used to name devices found within the cloud account within the LM portal
    + `disableStopTerminateHostMonitor` - If monitoring of stopped/terminated hosts is disabled
    + `disableTerminatedHostAlerting` - If alerting should be disabled when a cloud device is terminated
    + `monitoringRegionInfos` - The regions this group will monitor (array of strings)
    + `monitoringRegions` - The regions this group will monitor (array of strings) (deprecated) (computed)
    + `nameFilter` - Filter to determine which cloud devices should be monitored based on the device name (array of strings)
    + `normalCollectorConfig` - Configuration of 'normal' collector installed to the cloud (ie an EC2 instance)
      + `appliesTo` - AppliesTo expression to match or filter device
      + `autoBalancedCollectorGroupId` - Auto balanced collector group id assigned to ec2 or azure vm (int)
      + `collectorId` - Normal collector ID assigned to ec2 or azure vm (int)
      + `priority` - Collector match priority, smaller value has higher priority (int)
      + `usePublicIP` - Use cloud device public ip or not (boolean)
    + `selectAll` - Whether or not to use all regions (can be used instead of monitoringRegions and monitoringRegionInfos)
    + `tags` - Tags used to filter whether or not a cloud device is included in this group (array of CloudTagFilters)
      + `name` - The name of the device property that this filter will appy to (i.e. system.aws.tag.os). Case sensitive for AWS resources
      + `operation` - The operation that determines how this filter will be applied (options are 'exclude' or 'include')
      + `value` - The glob expression that the filter will match on
    + `useDefault` - Whether or not to use the default settings for this service. When set to true in a Service's settings, all other fields will be ignored (required)
  + `devices` - Cloud devices to monitor in the group (computed)
    + `deviceType` - Cloud device type (2 for AWS Device) (int)
    + `requiredProps` - Required device properties (array of strings)
  + `services` - Cloud account services to monitor. This is an object with keys for each AWS/Azure service that have the same arguments as `default`. See example above (all service keys alternate letters and underscores, but not numbers; SQS=s_q_s, SAGEMAKER=s_a_g_e_m_a_k_e_r, EC2=e_c2 etc.)
* `group_type` - The type of device group: normal and dynamic device groups will have groupType=Normal, and AWS groups will have a groupType value of AWS/SERVICE (e.g. AWS/AwsRoot, AWS/S3, Azure/AzureRoot etc.)
* `parent_id` - The id of the parent group for this device group (the root device group has an Id of 1)

## Import

device groups can be imported using their device group ID or full path
```
$ terraform import logicmonitor_device_group.my_device_group 66
$ terraform import logicmonitor_device_group.my_device_group Production/SBA/Linux
```