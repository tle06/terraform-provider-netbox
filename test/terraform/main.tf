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
  name = "mysite"
  slug = trimspace(lower(replace("mysite"," ","-")))
  region_id = 12
  status = "planned"
  tenant_id = 6
  facility = "tf facility"
  asn_id = 65000
  time_zone = "Africa/Asmara"
  description = "My super site"
  physical_address = "my physicla address"
  shipping_address = "my shipping address"
  latitude = "10.800000"
  longitude = "11.600000"
  contact_name = "John doe"
  contact_phone = "+33 7 45 81 81 93"
  #contact_email = "john.doe@gmail.com" # not working
  comments = "this is a comment"
  tags  {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tag-one.slug
    }
  tags {
    name = netbox_tag.tag-two.name
    slug = netbox_tag.tag-two.slug
    }
    
  custom_fields = {
    tf-test = "customFieldValue"
  }
}

resource "netbox_dcim_rack" "example" {
  name = "TF rack"
  site_id = netbox_dcim_site.example.id
  facility = "tf facility rack"
  tenant_id = 6
  status = "available" 
  role_id = 2
  serial = "test serial" 
  asset_tag = netbox_tag.tag-one.name
  type = "2-post-frame"
  width = 23
  u_height = 15
  desc_units = true
  outer_width = 11
  outer_depth = 10
  outer_unit = "mm"
  comments = "new comment"
  tags {
    name = netbox_tag.tag-one.name
    slug = netbox_tag.tag-one.slug
  }

  tags {
    name = netbox_tag.tag-two.name
    slug = netbox_tag.tag-two.slug
  }

  custom_fields = {
    rackCustomField = "rackCustomeFieldValue"
  }
}

resource "netbox_ipam_prefix" "example"{
  prefix = "10.0.0.0/16"
  site_id = netbox_dcim_site.example.id
  status = "active"
  
}


resource "netbox_dcim_device" "example" {
  device_type_id = 7
  device_role_id = 4
  site_id = netbox_dcim_site.example.id
  tenant_id = 6
  comments = "my comment"
  status = "active"
  asset_tag = netbox_tag.tag-one.name
  cluster_id = 9
  serial = "test serial"
  face = "front"
  name = "test device"
  # parent_device_id = 88
  platform_id = 2
  position_id = 1
  #primary_ip = "10.0.0.1/16"
  #primary_ip4_id = 9
  # primary_ip6_id =
  rack_id = netbox_dcim_rack.example.id
  # vc_position_id =
  # vc_priority_id =
  # virtual_chassis_id =
  tags {
    name = netbox_tag.tag-two.name
    slug = netbox_tag.tag-two.slug
  }
  custom_fields = {
    deviceCsutomField = "deviceCustomFieldValue"
  }
}