# netbox_ipam_prefix Resource

Creates a rack.

## Example Usage

```hcl
resource "netbox_tag" "tag-one" {
  name  = "tag one"
  slug  = "tag-one"
  color = "ff0000"
}

resource "netbox_dcim_site" "example" {
  name = "example"
  slug = "example"
}

resource "netbox_dcim_rack" "example" {
  name = "TF rack"
  site_id = netbox_dcim_site.example.id
}

resource "netbox_dcim_device" "example" {
  device_type_id = 7
  device_role_id = 4
  site_id = netbox_dcim_site.example.id

  tags {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tag-one.slug
  }

  custom_fields = {
    deviceCsutomField = "deviceCustomFieldValue"
  }
```

## Argument Reference

* `device_type_id` - (Required) The type ID of the device.

* `device_role_id` - (Required) The role ID of the device.

* `site_id` - (Required) The site ID to lcoate the device.

* `tenant_id` - (Optional) The tenant ID to add.

* `comments` - (Optional) The comment on the device.

* `status` - (Optional) The status of the device. Possible value: `offline`, `active`, `planned`, `staged`, `failed`, `inventory`, `decommissioning`. Default value is `active`.

* `asset_tag` - (Optional) Asset tag of the device.

* `cluster_id` - (Optional) The cluster ID to add.

* `serial` - (Optional) The serial of the device.

* `face` - (Optional) The rack face of the device. Possible value: `front`, `back`. Need `rack_id` value.

* `name` - (Optional) The name of the device.

* `platform_id` - (Optional) The platform ID of the device.

* `position_id` - (Optional) The position in the rack of the device. Need `rack_id` value.

* `rack_id` - (Optional) The rack ID of the device.

* `tags` - (Optional) List of tags to assign to the device. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the device. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
