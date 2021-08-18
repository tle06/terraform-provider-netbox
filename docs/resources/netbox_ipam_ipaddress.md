# netbox_ipam_prefix Resource

Creates an IP address.

## Example Usage

```hcl
resource "netbox_ipam_ipaddress" "example" {
  address = "10.0.0.1/16"
}
```

## Argument Reference

* `address` - (Required) The IP address to add.

* `nat_outside_id` - (Optional) the nat outside ID.

* `description` - (Optional) The description to add.

* `tenant_id` - (Optional) The tenant ID to add. Need a `vrf_id` to be set.

* `status` - (Optional) The status of the IP. Possible value: `active`, `reserved`, `deprecated`, `dhcp`, `slaac`. Default value is `active`.

* `role` - (Optional) The role of the IP address. Possible value: `loopback`, `secondary`, `anycast`, `vip`, `vrrp`, `hsrp`, `glbp`, `carp`.

* `vrf_id` - (Optional) The VRF of the IP address.

* `assigned_object_id` - (Optional) The assigned object ID to add.

* `assigned_object_type` - (Optional) The assgined object type to add.

* `dns_name` - (Optional) The DNS name to add.

* `nat_inside_id` - (Optional) The NAT inside ID to add.

* `tags` - (Optional) List of tags to assign to the IP address. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_extras_tag.example2.name
      slug = netbox_extras_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the IP address. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```
## Attribute Reference

* `id` - The prefix ID.
