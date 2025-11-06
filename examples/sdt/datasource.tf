resource "logicmonitor_sdt" "my_sdt"{
  type = "CollectorSDT"  # The values can be CollectorSDT | DeviceDataSourceInstanceSDT | DeviceBatchJobSDT | DeviceDataSourceInstanceGroupSDT | DeviceDataSourceSDT | DeviceEventSourceSDT | ResourceGroupSDT | ResourceSDT
  comment = "Scheduled maintenance window for system updates"
  sdt_type = "oneTime"
  start_date_time = 1767225600000  # November 30, 2025 00:00:00 UTC
  end_date_time = 1767232800000    # November 30, 2025 01:00:00 UTC  
  duration = 60
  timezone = "America/New_York"
  default_value = "2025-01-01T00:00:00Z"
  month_day = 7
  collector_id = "30329"
  #device_id = "3087872"
  #device_data_source_id = "25268320"
  #data_source_id = "552560846"
  #device_group_full_path = "DEVTS-21343 Dynamic Group"
  #device_group_id = "275527"
  #device_data_source_id = "24195849"
  #data_source_instance_id = "368309065"
  #device_data_source_instance_group_id = "27132393"
  #device_event_source_id = "2333036" 
  #device_batch_job_id = "64369"
}

data "logicmonitor_sdt" "my_sdt" {
       filter = "type~\"\"CollectorSDT"
 	   depends_on = [
	   logicmonitor_sdt.my_sdt
 	   ]
}

output "sdt" {
  description = "sdt"
  value       = data.logicmonitor_sdt.my_sdt.id
}

