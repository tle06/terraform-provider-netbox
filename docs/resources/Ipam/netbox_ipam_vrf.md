# netbox_ipam_prefix Resource

Creates a VRF.

## Example Usage

```hcl
resource "netbox_ipam_vrf" "example" {
  name = "example"
}
```

## Argument Reference

* `name` - (Required) The name of the VRF.
* `description` - (Optional) The descriptino to add.
* `tenant_id` - (Optional) The tenant ID to add.
* `enforce_unique` - (Optional) Enforce the unique Ip space. Possible value: `true`, `false`. Default value is `true`.
* `rd` - (Optional) The route distinguisher (RFC 4364) to add.

* `tags` - (Optional) List of tags to assign to the VRF. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the VRF. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
