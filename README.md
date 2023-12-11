# Terraform Provider for LogicMonitor

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://www.datocms-assets.com/2885/1629941242-logo-terraform-main.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.2.0
-	[Go](https://golang.org/doc/install) 1.20.3 (to build the provider plugin)

## Steps to downgrade or upgrade go version

To downgrade or upgrade Go to a specific version, you need to follow these steps:

Uninstall the current version of Go from your system by removing the Go installation directory.
Download the specific Go version binary for your operating system and architecture from the official website. You can find the previous releases of Go at:
```sh
 https://golang.org/dl/
 ```
Extract the downloaded archive to a location on your system.
Set the GOROOT environment variable to point to the extracted Go directory (the path where you extracted the Go binary).
Update your PATH environment variable to include the bin directory inside the extracted Go directory.
 ```sh
export GOROOT=/path/to/extracted/go
export PATH=$GOROOT/bin:$PATH
```
## Steps to downgrade or upgrade terraform version

To upgrade Terraform manually, follow these steps:

Visit the official HashiCorp releases page for Terraform at
```sh
 https://releases.hashicorp.com/terraform/
 ```
Find the desired version and click on it to see available downloads.
Download the appropriate archive for your operating system (e.g., ZIP for windows, tar.gz for macOS/Linux).
Extract the downloaded archive to a directory included in your system's PATH environment variable.
Verify the installation by running terraform version in the terminal.

To downgrade to a specific version manually, follow these steps:

Uninstall the currently installed Terraform version.
Follow the manual installation steps above, ensuring that you download and extract the desired version.
Remember to update any configuration files or scripts that rely on specific Terraform versions when upgrading or downgrading.

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

## Initialize the Terraform project

Run the following command to initialize the project and download necessary provider plugins:
```sh
terraform init
```
## Plan the execution

```sh
terraform plan
```
This will show you a preview of the changes that Terraform will make based on your configuration. Review the plan carefully before proceeding to apply it.

## Apply the changes
 
If you are satisfied with the execution plan and ready to apply the changes, run the following command:
```sh
terraform apply
```
Terraform will prompt you to confirm the execution plan one more time. Type yes and press Enter to proceed.
The Terraform execution will begin and create/update the necessary resources according to your configuration.

## Internal Documentation
```sh
https://confluence.logicmonitor.com/display/ENG/Development+Guide
```
