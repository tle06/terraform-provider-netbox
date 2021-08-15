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
  name = "test tf provider"
  slug = trimspace(lower(replace("test tf provider"," ","-")))
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