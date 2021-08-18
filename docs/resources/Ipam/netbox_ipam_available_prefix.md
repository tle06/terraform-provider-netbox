# netbox_ipam_available_prefix Resource

Creates an available child prefix within a parent.

## Example Usage

```hcl
resource "netbox_ipam_available_prefix" "example" {
  prefix_id     = 123
  prefix_length = 16
}
```

## Argument Reference

* `prefix_id` - (Required) The prefix ID.

* `prefix_length` - (Required) The number of bits of the prefix.

## Attribute Reference

* `prefix` - The address prefix.

* `family` - A value for the address family. Possible values are: `4` (IPv4) and `6` (IPv6).
