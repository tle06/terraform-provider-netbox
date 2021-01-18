# netbox_ipam_prefixes Data Source

Use this data source to get information about all prefixes.

## Example Usage

```hcl
data "netbox_ipam_prefixes" "example" {
  status  = "available"
  is_pool = true
}
```

## Argument Reference

* `contains` - (Optional) A case insensitive `contains` value, used to filter the results.

* `mask_length` - (Optional) A mask length used to filter the results.

* `prefix` - (Optional) - A prefix used to filter the results.

* `is_pool` - (Optional) - Whether the prefix is a pool.

* `region` - (Optional) - The name of a region.

* `role` - (Optional) - The name of a role.

* `site` - (Optional) - The name of a site.

* `status` - (Optional) - Operational status of the prefixes. Possible values are: `active`, `container`, `deprecated` and `reserved`.

* `tag` - (Optional) - The name of a tag.

* `tenant` - (Optional) - The name of a tenant.

* `family` - (Optional) - A value for the address family. Possible values include: `4` and `6`.

* `vrf_id` - (Optional) - The ID of the VRF.

* `within` - (Optional) - A case insensitive `within` value, used to filter the results.

* `within_include` - (Optional) A case insensitive `within_include` value, used to filter the results.

## Attribute Reference

* `results` - One or more `results` blocks as defined below.

The `results` block contains:

* `id` - The ID of the prefix.

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
