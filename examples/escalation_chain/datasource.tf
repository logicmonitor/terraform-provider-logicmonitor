resource "logicmonitor_escalation_chain" "my_escalation_chain"{
  throttling_alerts = 40
  enable_throttling = true
  destinations = [
    {
      period = [{
            week_days = 2
            timezone = "UTC"
            start_minutes = 10
            end_minutes   = 15
        }],
        stages = [{
            type    = "Admin"
            addr    = "unicornsparkles@rainbow.io"
            method  = "EMAIL"
            contact = "78362637"
        }]
        type = "timebased"
    }
]
  name = "Escalation Chain Test"
  description = "escalation chain test"
  cc_destinations = [
    {
      method = "EMAIL"
      contact = "string"
      type = "Admin"
      addr = "unicornsparkles@rainbow.io"
    }
  ]
  throttling_period = 30
}

data "logicmonitor_escalation_chain" "my_escalation_chain" {
    filter = "description~\"escalation chain test\""
 	depends_on = [
		logicmonitor_escalation_chain.my_escalation_chain
 	]
}

output "escalationChain" {
  description = "escalation chain"
  value       = data.logicmonitor_escalation_chain.my_escalation_chain.id
}

