resource "logicmonitor_role" "my_role" {
  name        = "TF Role"
  description = "Full access with all privileges"
  two_f_a_required = true
  require_e_u_l_a  = true
  custom_help_label = "Internal Admin Support"
  custom_help_url   = "https://internal-docs.company.com/admin-help"
  role_group_id = 1
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
}
data "logicmonitor_role" "my_role" {
   filter = "description~\"role test\""
 	depends_on = [
		logicmonitor_role.my_role
 	]
}

output "role" {
  description = "role"
  value       = data.logicmonitor_role.my_role.id
}
