# netbox_ipam_prefix Data Source

Use this data source to get information about a prefix.

## Example Usage

```hcl
data "netbox_ipam_prefix" "example" {
  prefix_id = 123
}
```

## Argument Reference

* `prefix_id` - (Required) The ID of the prefix.

## Attribute Reference

* `prefix` - The address prefix.

* `description` - A description for the prefix.

* `family` - A `family` block as defined below.

* `site` - A `site` block as defined below.

* `vrf` - A `vrf` block as defined below. Represents a virtual routing and forwarding (VRF) domain.

* `tenant` - A `tenant` block as defined below.

* `vlan` - A `vlan` block as defined below.

* `status` - A `status` block as defined below.

* `role` - A `role` block as defined below.

* `is_pool` - Whether this prefix is a pool. All IP addresses within this prefix are considered usable.

* `tags` - A list of tags for the prefix.

* `custom_fields` - A mapping of custom fields for the prefix.

The `family` block contains:

* `value` - A value for the address family. Possible values are: `4` and `6`.

* `label` - A label for the address family. Possible values are: `IPv4` and `IPv6`.

The `site` block contains:

* `id` - The ID of the site.

* `name` - The name of the site.

* `slug` - The site slug.

The `vrf` block contains:

* `id` - The ID of the VRF.

* `name` - The name of the VRF.

* `rd` - The route distinguisher (RD).

The `tenant` block contains:

* `id` - The ID of the tenant.

* `name` - The name of the tenant.

* `slug` - The tenant slug.

The `vlan` block contains:

* `id` - The ID of the Netbox VLAN object.

* `vid` - The configured VLAN ID.

* `name` - The configured VLAN name.

* `display_name` - The display name of the VLAN.

The `status` block contains:

* `value` - A value for the operational status of the prefix. Possible values are: `active`, `container`, `deprecated` and `reserved`.

* `label` - A label for the operational status of the prefix. Possible values are: `Active`, `Container`, `Deprecated` and `Reserved`.

The `role` block contains:

* `id` - The ID of the role.

* `name` - The name of the role.

* `slug` - The role slug.
