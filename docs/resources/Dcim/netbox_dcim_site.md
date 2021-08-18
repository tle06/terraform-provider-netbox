# netbox_ipam_prefix Resource

Creates a site.

## Example Usage

```hcl
resource "netbox_tag" "tag-one" {
  name  = "tag one"
  slug  = "tag-one"
  color = "ff0000"
}

resource "netbox_tag" "tag-two" {
  name  = "tag two"
  slug  = "tag-two"
  color = "ff0000"
}

resource "netbox_dcim_site" "example" {
  name = "example"
  slug = "example"
  
  tags  {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tage-one.slug
    }

  tags {
    name = netbox_tag.tag-tow.name
    slug = netbox_tag.tag-two.slug
    }

  custom_fields = {
    customFieldName = "customFieldValue"
  }
```

## Argument Reference

* `name` - (Required) The name to add.
  
* `slud` = (required) the slug of the site to add
  
* `region_id` - (Optional) The ID of the region to assign to the site.
  
* `status` - (Optional) Operational status of this site. Possible values are: `planned`, `staging`, `active`, `decommissioning`, `retired`. Default value: `active`.
  
* `tenant_id` - (Optional) The ID of a tenant to assign to the site.
  
* `facility` - (Optional) A facility for the site.

* `asn_id` - (Optional) The BGP ASN for the site (eg. `65000`)

* `time_zone` - (Optional) A time_zone for the site. Value are [TZ_Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) (eg. `Africa/Abidjan`)
  
* `physical_address` - (Optional) The physical address for the site.

* `shipping_address` - (Optional) The shipping address for the site.

* `latitude` - (Optional) The latitude for the site (eg .`10.800000`). To configure input 2 digit then a dot then 6 digits.

* `longitude` - (Optional) The longitute for the site (eg .`10.800000`). To configure input 2 digit then a dot then 6 digits.

* `contact_name` - (Optional) The contact name for the site.

* `contact_phone` - (Optional) The contatc phone for the site.

* `comments` - (Optional) A comments for the site.
  
* `tags` - (Optional) List of tags to assign to the Site. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the site. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
