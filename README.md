# Terraform Provider for LogicMonitor

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/source/assets/images/logo-text.svg" width="600px">

## Requirements


-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

## Testing The Provider

Clone repository to: `$GOPATH/src/github.com/logicmonitor/terraform-provider-logicmonitor`

```sh
$ mkdir -p $GOPATH/src/github.com/logicmonitor; cd $GOPATH/src/github.com/logicmonitor
$ git clone git@github.com:logicmonitor/terraform-provider-logicmonitor
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/logicmonitor/terraform-provider-logicmonitor
$ make build
```

Copy the provider to your terraform project directory

0.9.x:
```sh
$ cp $GOPATH/bin/terraform-provider-logicmonitor /path/to/terraform/project
```

0.10.x:
```sh
$ mkdir -p /terraform-project-path/terraform.d/plugins/darwin_amd64/
$ cp $GOPATH/bin/terraform-provider-logicmonitor /terraform-project-path/terraform.d/plugins/darwin_amd64/
```

## Using the provider

LogicMonitor Terraform Provider has two methods for setting required arguments:
Environment Variables
```sh
export LM_API_ID=xyz
export LM_API_KEY=xyz
export LM_COMPANY=xyz
```
Terraform variables
```sh
# Configure the LogicMonitor provider
provider "logicmonitor" {
  api_id = "${var.logicmonitor_api_id}"
  api_key = "${var.logicmonitor_api_key}"
  company = "${var.logicmonitor_company}"
}
```

### Examples

This is an example snippet we added with one of our sample modules (ec2_host)
```
resource "logicmonitor_device" host{
  count    = "${var.count}"
  ip_addr = "${element(aws_instance.ec2_host.*.private_ip, count.index)}"
  displayname = "${element(aws_route53_record.ec2_host_private_dns.*.fqdn, count.index)}"
  collector = 123
}
```

And our terraform plan shows
```
+ module.ec2_host.logicmonitor_device.host.0
    collector:       "123"
    disablealerting: "true"
    ip_addr:         "${element(aws_instance.ec2_host.*.private_ip, count.index)}"

+ module.ec2_host.logicmonitor_device.host.1
    collector:       "123"
    disablealerting: "true"
    ip_addr:         "${element(aws_instance.ec2_host.*.private_ip, count.index)}"
```


### Data Providers
Using our data provider, we can add filters for things like host group or collector lookups.
```
# Find a device group containing name of Production with select custom properties
data "logicmonitor_finddevicegroups" "devicegroups" {

  filters {
    "property" = "name"
    "operator" = "~"
    "value" = "Production"
  },

  filters {
    "CustomPropertyName" = "app"
    "operator" = ":"
    "CustomPropertyValue" = "test"
  }
}
```

Each data instance will export one or more attributes, which can be interpolated into other resources using variables of the form data.TYPE.NAME.ATTR. For example:
```
# in this case data.logicmonitor_finddevicegroups.devicegroups.id will return a string of comma separated ids.
resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disablealerting = true
  hostgroupid = "${data.logicmonitor_finddevicegroups.devicegroups.id}"
}
```

Here is another example
```
# Find the most recent collector created that is online and with hostname containing the string "test"
data "logicmonitor_findcollectors" "collectors" {
  filters {
    "property" = "hostname"
    "operator" = "~"
    "value" = "test"
  },
"most_recent" = true
}
```

## Supported Resources
### Add Device
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip_addr** | **string** | set device ip address | [required]
**displayname** | **string** | set device displayname | [optional] [default to ip_addr if field null ]
**disablealerting** | **bool** | alerting status on device | [optional] [default to true]
**collector** | **int32** | set preferred collector for device | [required]
**hostgroupid** | **string** | set device parent host group based on hostgroupid | [optional] [default to null]
**hostgroupname** | **string** | set device parent host group based on hostgroupname  | [optional] [full name of group required, will be overwritten if hostgroupid is not null]
**description** | **string** | set description of device | [optional] [default to null]
**properties** | **map** | set device properties | [optional] [default to null, see sample main.tf]

### Add Device Group
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | name of device group | [required]
**parentId** | **string** | set device group parentId | [optional] [default to null ]
**disablealerting** | **bool** | alerting status on device group| [optional] [default to true]
**description** | **string** | set description of device group | [optional] [default to null]
**properties** | **map** | set device group properties | [optional] [default to null, see sample main.tf]
**fullpath** | **string** | set full path of device group | [optional] [default to null]

### Add Collector Group
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | name of collector group | [optional] [default to null]
**description** | **string** | set description of collector group | [optional] [default to null]

## Supported Data Sources
### Collector lookup
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**filters** | **map** | filters for collectors | [optional] [default to null] [property (string), operator(string), and value(string) must all be completed for filter to work. See help.logicmonitor.com for appropriate filters]
**most_recent** | **bool** | find latest collector added to portal | [optional] [default to false]
**size** | **int** | The number of results to display. Max is 1000. | [optional] [default to 50]
**offset** | **int** | The number of results to offset the displayed results by | [optional] [default to 0]

### HostGroup lookup
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**filters** | **map** | filters for collectors | [optional] [default to null] [property (string), operator(string), and value(string) must all be completed for filter to work, the same goes for CustomPropertyName (string), operator(string), and CustomPropertyValue(string). See help.logicmonitor.com for appropriate filters]
**size** | **int** | The number of results to display. Max is 1000. | [optional] [default to 50]
**offset** | **int** | The number of results to offset the displayed results by | [optional] [default to 0]

## Developing the Provider

### Contributing Resources

### Development Environment

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-logicmonitor
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
