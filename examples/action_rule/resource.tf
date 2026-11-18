resource "logicmonitor_action_rule" "my_action_rule" {
	name            = "Action_rule_test"
	action_chain_id = 216
	level_str       = "All"
	datasource      = "Pingds05"
	instance        = "*"
	datapoint       = "*"
	device_groups   = ["Parag_device_group"]
	devices         = ["*"]
}