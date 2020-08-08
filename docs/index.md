# Netbox Provider

Use this provider to manage [Netbox](https://netbox.readthedocs.io/) resources.

## Example Usage

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

## Argument Reference

* `host` - (Required) The Netbox hostname to connect to. It can also be sourced from the `NETBOX_HOST` environment variable.

* `token` - (Optional) The API token used to authenticate with Netbox. It can also be sourced from the `NETBOK_TOKEN` environment variable.
