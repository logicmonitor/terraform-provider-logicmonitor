resource "logicmonitor_action_chain" "my_action_chain" {
	name = "Action_chain_test"
	stages {
		id   = 625
		type = "diagnosticSource"
	}
}
