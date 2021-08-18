# netbox_ipam_prefix Resource

Creates a prefix.

## Example Usage

```hcl
resource "netbox_ipam_prefix" "example" {
  prefix = "10.0.0.0/16"
  status = "reserved"
}
```

## Argument Reference

* `prefix` - (Required) The prefix to add. This should be IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. `10.0.0.0/16`).

* `description` - (Optional) A description for the prefix.

* `status` - (Optional) Operational status of this prefix. Possible values are: `active`, `container`, `deprecated` and `reserved`.

* `site_id` - (Optional) The ID of a site to assign to the prefix.

* `vrf_id` - (Optional) The ID of a VRF to assign to the prefix.

* `tenant_id` - (Optional) The ID of a tenant to assign to the prefix.

* `vlan_id` - (Optional) The ID of a VLAN to assign to the prefix.

* `role_id` - (Optional) The ID of a role to assign to the prefix.

* `is_pool` - (Optional) Whether this prefix is a pool. All IP addresses within this prefix are considered usable.

* `tags` - (Optional) List of tags to assign to the prefix.

* `custom_fields` - (Optional) A mapping of custom fields to assign to the prefix.

## Attribute Reference

* `id` - The prefix ID.

* `family` - A value for the address family. Possible values are: `4` (IPv4) and `6` (IPv6).
