resource "logicmonitor_device_group" "my_device_group" {
	description = "normal device group test" 
	disable_alerting = true
	enable_netflow = false
	default_collector_id = 10
	name = "meng test device group test"
	parent_id = 1
	custom_properties { 
          name = "addr"      
          value = "127.0.0.1" 
     }
	group_type = "Normal"
}


data "logicmonitor_device_group" "my_device_group" {
    filter = "description~\"normal device group test\""
	depends_on = [
		logicmonitor_device_group.my_device_group
	]
}

output "deviceGroup" {
  description = "devices group list"
  value       = data.logicmonitor_device_group.my_device_group.id
}

/* AWS Device Group Testing */
/*
provider "aws" {
  access_key = ""
  secret_key = ""
  region     = ""
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
//*/

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