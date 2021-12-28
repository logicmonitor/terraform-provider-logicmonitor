# Terraform Provider for LogicMonitor

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://www.datocms-assets.com/2885/1629941242-logo-terraform-main.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.14.x
-	[Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

## Building the Provider (Internal)
Clone repository:
```sh
$ git clone ssh://git@stash.logicmonitor.com:7999/teamintegrations/terraform-integration.git
```
Enter the provider directory and build the provider:
```sh
$ cd terraform-integration/terraform-provider-logicmonitor/
$ make
```
The make file will then generate the code, build the binary, and copy it to the Terraform plugin directory. 

## Using the provider

The LogicMonitor Terraform Provider has two methods for setting required arguments:
Environment Variables
```sh
export LM_API_ID=xyz
export LM_API_KEY=xyz
export LM_COMPANY=xyz
```

Provider Initialization
```sh
provider "logicmonitor" {
  api_id = var.logicmonitor_api_id
  api_key = var.logicmonitor_api_key
  company = var.logicmonitor_company
}
```
Example usage can be found in the /terraform-integration/examples directory.

## Internal Documentation
```sh
https://confluence.logicmonitor.com/display/ENG/Development+Guide
```
