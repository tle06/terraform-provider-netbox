resource "netbox_extras_tag" "tag-one" {
  name  = "tag one"
  slug  = "tag-one"
  color = "ff0000"
}

resource "netbox_extras_tag" "tag-two" {
  name  = "tag two"
  slug  = "tag-two"
  color = "ff0000"
}

resource "netbox_tenancy_tenant" "example" {
  name = "test terraform"
  slug = "test-terraform"
}

resource "netbox_dcim_region" "root-region" {
  name = "terraform root region"
  slug = "terraform-root-region"
  description = "description for terraform"
}

resource "netbox_dcim_region" "child-region" {
  name = "terarform child region"
  slug = "terarform-child-region"
  parent_id = netbox_dcim_region.root-region.id
  description = "description for terraform"
}

resource "netbox_dcim_site" "example" {
  name = "mysite"
  slug = trimspace(lower(replace("mysite"," ","-")))
  region_id = netbox_dcim_region.child-region.id
  status = "planned"
  tenant_id = netbox_tenancy_tenant.example.id
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
  contact_email = "john.doe@gmail.com" # not working
  comments = "this is a comment"
  tags  {
    name = netbox_extras_tag.tag-one.name
    slug = netbox_extras_tag.tag-one.slug
    }
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
    }
    
}

resource "netbox_dcim_rack" "example" {
  name = "terraform rack"
  site_id = netbox_dcim_site.example.id
  facility = "tf facility rack"
  tenant_id = netbox_tenancy_tenant.example.id
  status = "available" 
  role_id = 2
  serial = "test serial" 
  asset_tag = netbox_extras_tag.tag-one.name
  type = "2-post-frame"
  width = 23
  u_height = 15
  desc_units = true
  outer_width = 11
  outer_depth = 10
  outer_unit = "mm"
  comments = "new comment"
  tags {
    name = netbox_extras_tag.tag-one.name
    slug = netbox_extras_tag.tag-one.slug
  }

  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
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
  tenant_id = netbox_tenancy_tenant.example.id
  comments = "my comment"
  status = "active"
  asset_tag = netbox_extras_tag.tag-one.name
  cluster_id = 9
  serial = "test serial"
  face = "front"
  name = "test device"
  platform_id = 2
  position_id = 1
  rack_id = netbox_dcim_rack.example.id
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
  }

}


resource "netbox_ipam_vlan" "tagged-vlan" {
  name = "test terraform"
	vid = 300
	site_id = netbox_dcim_site.example.id
	tenant_id = netbox_tenancy_tenant.example.id
	status = "reserved"
	role_id = 1
  description = "test terraform"
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
  }
}

resource "netbox_ipam_vlan" "untagged-vlan" {
  name = "test terraform"
	vid = 400
	site_id = netbox_dcim_site.example.id
	tenant_id = netbox_tenancy_tenant.example.id
	status = "reserved"
	role_id = 1
  description = "test terraform"
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
  }
}



resource "netbox_dcim_interface" "example" {

  device_id = netbox_dcim_device.example.id
  type = "virtual"
  name = "test interface"
  tagged_vlan = [netbox_ipam_vlan.tagged-vlan.id]
  connection_status = true
  enabled = true
  management_only = true
  label = "label"
  mac_address = "00:00:00:00:00:00"
  mode = "access"
  description = "test"
  untagged_vlan_id = netbox_ipam_vlan.untagged-vlan.id
  mtu = 1000
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
  }
}


# resource "netbox_circuits_circuit" "example" {

#   cid ="test cid"
#   type_id = 1
#   provider_id = 2
#   #status
#   # tenant_id
#   # commit_rate
#   # comments
#   # description
#   # install_date
#   # tags
#   # custom_fields

# }


resource "netbox_ipam_vrf" "example" {
  name = "terraform vrf"
  # description = "test description"
  # tenant_id = netbox_tenancy_tenant.example.id
  # enforce_unique = false
  # rd = "64512:900:192.168"
  # custom_fields = {}
  # tags {
  #   name = netbox_extras_tag.tag-two.name
  #   slug = netbox_extras_tag.tag-two.slug
  # }

}

resource "netbox_ipam_ipaddress" "example" {
  address = "10.0.0.1/16"
  
  description = "test"
  tenant_id = netbox_tenancy_tenant.example.id
  status = "reserved"
  role = "vip"
  vrf_id = netbox_ipam_vrf.example.id
  # nat_outside_id = 64
  # assigned_object_id
  # assigned_object_type
  dns_name = "test.example.com"
  # nat_inside_id
  tags {
    name = netbox_extras_tag.tag-two.name
    slug = netbox_extras_tag.tag-two.slug
  }
  # custom_fields = {
  #   ipAddressCustomField = "ipAddressCustomFieldValue"
  # }
}


resource "netbox_circuits_provider" "example" {
  name = "example"
  slug = "example"
  # asn = "65000"
  # account = "my account"
  # admin_contact = "john doe"
  # comments = "comment"
  # noc_contact = "john doe 2"
  # portal_url = "https://demo.netbox.dev"
  # tags {
  #   name = netbox_extras_tag.tag-two.name
  #   slug = netbox_extras_tag.tag-two.slug
  # }
  # custom_fields = {
  #   customFieldProvider ="customFieldProviderValue"
  # }
}