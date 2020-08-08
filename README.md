# Terraform Provider for Netbox

![](https://github.com/innovationnorway/terraform-provider-netbox/workflows/test/badge.svg)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12
-	[Go](https://golang.org/doc/install) >= 1.14

## Usage

```hcl
provider "netbox" {
  host  = "http://localhost:8000"
  token = "66a48ac409ec56b3f345eee3d10a42fa2fc1b8b9"
}

resource "netbox_ipam_prefix" "example" {
  prefix = "10.0.0.0/16"
  status = "reserved"
}
```

## Contributing

To build the provider:

```sh
$ go build
```

To test the provider:

```sh
$ go test -v ./...
```

To run all acceptance tests:

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ TF_ACC=1 go test -v ./...
```

To run a subset of acceptance tests:

```sh
$ TF_ACC=1 go test -v ./... -run=TestAccIpamPrefix
```
