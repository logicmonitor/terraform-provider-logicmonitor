## 2.0.21 (May, 2025)

* Domain field added in provider block to handle gov domains
* Special handling for 'need_stc_grp_and_sorted_c_p' query param


## 2.0.20 (April 11, 2025)

* HostGroupIds and custom properties fix

## 2.0.19 (December 23, 2024)

* Updated implementation of Alert rule for allowing escalation interval of 0 
* Documentation update for custom properties

## 2.0.18 (October 03, 2024)

* Website resource auth block handling

## 2.0.17 (July 11, 2024)

* Auto Discovery Feature Added And Provider Documentation Update
* Use Default Root and Website schema parameter fix


## 2.0.16 (May 30, 2024)

* SSL expiration check added for website and datasource documentation update

## 2.0.15 (May 8, 2024)

* Test location parameter added for website resource
* Escalation chain resource stage field fix

## 2.0.14 (March 22, 2024)

* Datasource documentation update

## 2.0.13 (February 12, 2024)

* Service insights support added
* **New Resource:** `logicmonitor_datasource`

## 2.0.12 (December 15, 2023)

* Website resource fields added and documentation update

## 2.0.11 (October 9, 2023)

* **New Resource:** `logicmonitor_website_group`

## 2.0.10 (September 15, 2023)

* **New Resource:** `logicmonitor_website`

* Escaltion chain bug fix for week_days

## 2.0.9 (August 10, 2023)

* **New Resource:** `logicmonitor_alert_rule`

## 2.0.8 (August 1, 2023)

* Escalation chain documentation update

## 2.0.7 (August 1, 2023)

* **New Resource:** `logicmonitor_escalation_chain`

## 2.0.6 (April 26, 2023)

* Added support for the following services:

Azure/KeyVault
Azure/LoadBalancers
Azure/BackupProtectedItems
Azure/RecoveryServices
Azure/AutomationAccount
Azure/VirtualMachineScaleSet
Azure/LogAnalyticsWorkspaces
Azure/VirtualNetworkGateway
Azure/ServiceBus
Azure/DataFactory
Azure/ApplicationGateway
Azure/EventHub
Azure/VirtualDesktop
Azure/SQLElasticPool
Azure/TrafficManager
Azure/RedisCache
Azure/CosmosDB
Azure/PostgreSQL
Azure/RecoveryProtectedItems
Azure/ApiManagement
Azure/ExpressRouteCircuit
Azure/SQLManagedInstance
Azure/CognitiveServices
Azure/FrontDoors
Azure/Firewall
Azure/MySQL
Azure/SynapseWorkSpaces
Azure/CognitiveSearch

## 2.0.5 (February 22, 2023)

* In this release, we are adding functionality to onboard Azure Cloud Accounts in TF Provider.

## 2.0.4 (December 12, 2022)

* In this release license information added.

## 2.0.3 (September 26, 2022)
ENHANCEMENTS:
* Add experimental `bulk_resource` configuration field to the provider.

## 2.0.2 (May 28, 2022)

FEATURES:
* **New Resource:** `logicmonitor_dashboard` 
* **New Resource:** `logicmonitor_dashboard_group`
* **New DataSource:** `logicmonitor_collector` 
* **New DataSource:** `logicmonitor_collector_group` 
* **New Datasource:** `logicmonitor_device`
* **New Datasource:** `logicmonitor_device_group`
* **New Datasource:** `logicmonitor_dashboard`
* **New Datasource:** `logicmonitor_dashboard_group`

## 2.0.1 (May 6, 2022)

NOTES:

* This release adds documentation for provider as per 2.0.0 release.

## 2.0.0 (December 28, 2021)

NOTES:

* Implementation of our new LogicMonitor Provider

## 1.3.4 (March 16, 2021)

IMPROVEMENTS:

* Add possibility to create a service using the device resource

## 1.3.3 (January 29, 2021)

This is patch increment of the previous version and does not introduce any new functionality.

## 1.3.0 (February 03, 2020)

FEATURES:
* **New Resource:** `logicmonitor_dashboard` ([#22](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/22))
* **New Resource:** `logicmonitor_dashboard_group` ([#22](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/22))
* **Import Functionality:** `logicmonitor_dashboard` ([#22](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/22))
* **Import Functionality:** `logicmonitor_dashboard_group` ([#22](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/22))

## 1.2.1 (June 18, 2019)

IMPROVEMENTS:

* Add use LogicMonitor API v2 ([#16](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/16))
* Add support for Terraform v0.12.0 ([#18](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/18))

## 1.2.0 (April 30, 2018)

FEATURES:
* **Import Functionality:** `logicmonitor_device` ([#4](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/4))
* **Import Functionality:** `logicmonitor_device_group` ([#4](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/4))

## 1.1.0 (March 19, 2018)

FEATURES:
* **New Resource:** `logicmonitor_collector` ([#1](https://github.com/terraform-providers/terraform-provider-logicmonitor/issues/1))

## 1.0.0 (October 11, 2017)

NOTES:

* Initial implementation of the LogicMonitor Provider
