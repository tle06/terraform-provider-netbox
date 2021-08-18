# netbox_ipam_prefix Resource

Creates a provider.

## Example Usage

```hcl
resource "netbox_circuits_provider" "example" {
  name = "example"
  slug = "example"
}
```

## Argument Reference

* `name` - (Required) The name of the provider.

* `slug` - (Required) The slug to add.

* `asn` - (Optional) The ASN number.

* `account` - (Optional) The client account.

* `admin_contact` - (Optional) The admin contct to add.

* `comments` - (Optional) The comments to add.

* `noc_contact` - (Optional) The NOC contact to add.

* `portal_url` - (Optional) The provider portal URL.

* `tags` - (Optional) List of tags to assign to the provider. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the provider. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
