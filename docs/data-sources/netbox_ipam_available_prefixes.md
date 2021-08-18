# netbox_ipam_available_prefixes Data Source

Use this data source to get available child prefixes within a parent.

## Example Usage

```hcl
data "netbox_ipam_available_prefixes" "example" {
  prefix_id = 123
}
```

## Argument Reference

* `prefix_id` - (Required) The ID of the prefix.

## Attribute Reference

* `prefixes` - One or more `prefix` blocks as defined below.

The `prefix` block contains:

* `prefix` - The address prefix.

* `family` - A value which represents the address family. Possible values are: `4` (IPv4) and `6` (IPv6).

* `vrf` - A `vrf` block as defined below. Represents a virtual routing and forwarding (VRF) domain.

The `vrf` block contains:

* `id` - The ID of the VRF.

* `name` - The name of the VRF.

* `rd` - The route distinguisher (RD).
