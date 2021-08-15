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

  tags {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tag-one.slug
  }

  custom_fields = {
    rackCustomField = "rackCustomeFieldValue"
  }
```

## Argument Reference

* `name` - (Required) The name to add.
  
* `site_id` - (Required) The site to attache the rack.
  
* `facility` - (Optional) A facility for the rack.

* `tenant_id` - (Optional) The ID of a tenant to assign to the rack.

* `status` - (Optional) Status of this rack. Possible values are: `reserved`, `available`, `planned`, `active`, `deprecated`. Default value `active`

* `role_id` - (Optional) The role ID for teh rack.

* `serial` - (Optional) The serial of the rack.

* `asset_tag` - (Optional) Asseg tag to add.

* `type` - (Optional) Type for the rack. Possible values are: `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-cabinet`.

* `width` - (Optional) The width of the rack. Possible values are: `10`, `19`, `21`, `23`. Default value is `19`.

* `u_height` - (Optional) The Unit size of the rack. Default unit size `42`

* `desc_units` - (Optional) If the unit are ordered in descending. Possible value `true`, `false`. Default value `false`.

* `outer_width` - (Optional) The outer width of the rack. `outer_init` need to eb set.

* `outer_depth` - (Optional) The outer depth or the rack. `outer_init` need to eb set.

* `outer_unit` - (Optional) The outer unit. Possible value `mm`, `in`. Default value `mm`.

* `comments` - (Optional) A comments for the rack.

* `tags` - (Optional) List of tags to assign to the rack. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the rack. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
