# netbox_ipam_prefix Resource

Creates a region.

## Example Usage

```hcl
resource "netbox_tenancy_tenant" "example" {
  name = "example"
  slug = "example"
}

```

## Argument Reference

* `name` - (Required) The name to add.
  
* `slud` = (Required) the slug of the tenant to add
  
* `group_id` - (Optional) The ID of the group.

* `description` - (Optional) The description to add.
  
* `comments` - (Optional) The comment to add.
  
* `tags` - (Optional) List of tags to assign to the Tenant. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the Tenant. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  
## Attribute Reference

* `id` - The prefix ID.
