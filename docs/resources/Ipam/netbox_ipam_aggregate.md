# netbox_ipam_aggregate Resource

Creates a aggregate.

## Example Usage

```hcl
resource "netbox_ipam_aggregate" "example" {
  prefix = "10.0.0.0/16"
  rir_id = 2
}
```

## Argument Reference

* `prefix` - (Required) The prefix to add. This should be IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. `10.0.0.0/16`).

* `rir_id` - (Optional) The ID of a RIR to assign to the prefix.

## Attribute Reference

* `rir_id` - The rir ID.

* `family` - A value for the address family. Possible values are: `4` (IPv4) and `6` (IPv6).
