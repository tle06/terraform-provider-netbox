# netbox_ipam_prefix Resource

Creates an interface on a device.

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


resource "netbox_dcim_device" "example" {
  device_type_id = 7
  device_role_id = 4
  site_id = netbox_dcim_site.example.id
}

resource "netbox_dcim_interface" "example" {
  device_id = netbox_dcim_device.example.id
  type = "virtual"
  name = "example"
  tagged_vlan = [64]
  
  tags {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tag-one.slug
  }
}
```

## Argument Reference

* `device_id` - (Required) The device ID where the interface need to be link.

* `type` - (Required) The interface type. Possible value: `virtual`, `lag`, `100base-tx`, `1000base-t`, `2.5gbase-t`, `5gbase-t`, `10gbase-t`, `10gbase-cx4`, `1000base-x-gbic`, `1000base-x-sfp`, `10gbase-x-sfpp`, `10gbase-x-xfp`, `10gbase-x-xenpak`, `10gbase-x-x2`, `25gbase-x-sfp28`, `40gbase-x-qsfpp`, `50gbase-x-sfp28`, `100gbase-x-cfp`, `100gbase-x-cfp2`, `200gbase-x-cfp2`, `100gbase-x-cfp4`, `100gbase-x-cpak`, `100gbase-x-qsfp28`, `200gbase-x-qsfp56`, `400gbase-x-qsfpdd`, `400gbase-x-osfp`, `ieee802.11a`, `ieee802.11g`, `ieee802.11n`, `ieee802.11ac`, `ieee802.11ad`, `ieee802.11ax`, `gsm`, `cdma`, `lte`, `sonet-oc3`, `sonet-oc12`, `sonet-oc48`, `sonet-oc192`, `sonet-oc768`, `sonet-oc1920`, `sonet-oc3840`, `1gfc-sfp`, `2gfc-sfp`, `4gfc-sfp`, `8gfc-sfpp`, `16gfc-sfpp`, `32gfc-sfp28`, `128gfc-sfp28`, `infiniband-sdr`, `infiniband-ddr`, `infiniband-qdr`, `infiniband-fdr10`, `infiniband-fdr`, `infiniband-edr`, `infiniband-hdr`, `infiniband-ndr`, `infiniband-xdr`, `t1`, `e1`, `t3`, `e3`, `cisco-stackwise`, `cisco-stackwise-plus`, `cisco-flexstack`, `cisco-flexstack-plus`, `juniper-vcp`, `extreme-summitstack`, `extreme-summitstack-128`, `extreme-summitstack-256`, `extreme-summitstack-512`, `other`.

* `name` - (Required) The name of the interface.

* `tagged_vlan` - (Required) A list of VLAN ID that will be bind to the interface (ei. `[1,2]`). VLAN need to belong to the site or being global.

* `connection_status` - (Optional) The status of the interface. Possible value: `true`, `false`.

* `enabled` - (Optional) Indicate if the interface is enabled. Possible value: `true`, `false`. Default value is `true`

* `management_only` - (Optional) Indicate if the interface is dedciated to management. Possible value: `true`, `false`.

* `label` - (Optional) The label of the interface

* `mac_address` - (Optional) The mac address of the interface. Format should be `00:00:00:00:00:00`.

* `mode` - (Optional) The interface mode. Possible value: `access`, `tagged`, `tagged-all`

* `description` - (Optional) The description to add.

* `untagged_vlan_id` - (Optional) The untagged VLAN ID.

* `mtu` - (Optional) The MTU of the interface.

* `tags` - (Optional) List of tags to assign to the interface. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_tag.example2.name
      slug = netbox_tag.example2.slug
    }
  ```

## Attribute Reference

* `id` - The prefix ID.
