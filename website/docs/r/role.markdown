---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_role"
sidebar_current: "docs-logicmonitor-resource-role"
description: |-
  Provides a LogicMonitor role resource. This can be used to create and manage LogicMonitor roles.
---

# logicmonitor_role

Provides a LogicMonitor role resource. This can be used to create and manage LogicMonitor roles.

## Example Usage
```hcl
# Create a LogicMonitor role
resource "logicmonitor_role" "my_role" {
  custom_help_label = "Internal Support Resources"
  custom_help_url = "https://logicmonitor.com/support"
  description = "Administrator can do everything, including security-sensitive actions"
  name = "administrator"
  privileges {
    object_type = "dashboard_group"
    object_id   = "14519"  
    operation   = "write"
  }
  privileges {
    object_type = "dashboard"
    object_id   = "68605"  
    operation   = "write"
  }
  require_e_u_l_a = true
  role_group_id = 3
  two_f_a_required = true
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the role
   (string)
* `privileges` - The account privileges associated with the role. Privileges can be added to a role for each area of your account
   ([]*Privilege)
  + `objectId` - The privilege object identifier yes
  + `objectType` - The privilege object type. The values can be dashboard_group|dashboard|host_group|service_group|website_group|report_group|remoteSession|chat|setting|device_dashboard|help|logs|configNeedDeviceManagePermission|map|resourceMapTab|tracesManageTab|costOptimization|dexda|lmSupportAccess
  + `operation` - The privilege operation

The following arguments are **optional**:
* `custom_help_label` - The label for the custom help URL as it will appear in the 'Help & Support' dropdown menu (string)
* `custom_help_url` - The URL that should be added to the 'Help & Support' dropdown menu (string)
* `description` - The description of the role (string)
* `require_e_u_l_a` - Whether or not users assigned this role should be required to acknowledge the EULA (end user license agreement) (bool)
* `role_group_id` - The group Id of the role (int32)
* `two_f_a_required` - Whether Two-Factor Authentication should be required for this role (bool)

## Import

roles can be imported using their role ID or name
```
$ terraform import logicmonitor_role.my_role 66
$ terraform import logicmonitor_role.my_role administrator
```