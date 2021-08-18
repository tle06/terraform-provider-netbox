# netbox_ipam_prefix Resource

Creates a vlan.

## Example Usage

```hcl
resource "netbox_ipam_vlan" "example" {
  name = "example"
	vid = 300
}
```

## Argument Reference

* `name` - (Required) The name to add.

* `vid` - (Required) The VLAN ID to add.

* `site_id` - (Optional) The site ID where the VLAN belong.

* `tenant_id` - (Optional) The tenant ID to add.

* `status` - (Optional) The status of the VLNA. Possible value: `active`, `deprecated`,`reserved`. Default value is `active`.

* `role_id` - (Optional) The role ID of the VLAN.

* `description` - (Optional) The description to add.

* `tags` - (Optional) List of tags to assign to the vlan.

## Attribute Reference

* `id` - The prefix ID.
