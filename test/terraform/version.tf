# Configuration to test locally the provider
terraform {
  required_providers {
    netbox = {
      source = "tle06/netbox"
      version = "0.1.0-alpha.3"
    }
    netbox-local = {
        source = "terraform.example.com/local/netbox"
    }
  }
}