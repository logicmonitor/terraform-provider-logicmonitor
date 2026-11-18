resource "logicmonitor_action_chain" "my_action_chain" {
	name = "Action_chain_test"
	stages {
		id   = 625
		type = "diagnosticSource"
	}
}
data "logicmonitor_action_chain" "my_action_chain" {
    filter = "name~\"Action_chain_test\""
	depends_on = [
		logicmonitor_action_chain.my_action_chain
	]
}
output "actionChain" {
  description = "action chain"
  value       = data.logicmonitor_action_chain.my_action_chain.id
}