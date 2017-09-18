# Terraform Provider for LogicMonitor Testing

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
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-logicmonitor
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-logicmonitor
$ make build
```

## Using the provider

LogicMonitor Terraform Provider has two methods for setting required arguments:
Environment Variables
```sh
export LM_API_ID=xyz
export LM_API_KEY=xyz
export LM_COMPANY=xyz
```

Provider Initialization
```sh
provider "logicmonitor" {
  api_id = "${var.logicmonitor_api_id}"
  api_key = "${var.logicmonitor_api_key}"
  company = "${var.logicmonitor_company}"
}
```

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
