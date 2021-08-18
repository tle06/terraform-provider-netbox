# netbox_ipam_prefix Resource

Creates a region.

## Example Usage

```hcl
resource "netbox_dcim_region" "example" {
  name = "example"
  slug = "example"
}

```

## Argument Reference

* `name` - (Required) The name to add.
  
* `slud` = (Required) the slug of the site to add
  
* `parent_id` - (Optional) The ID of the parent region.

* `description` - (Optional) The description to add.
  
## Attribute Reference

* `id` - The prefix ID.
