# Configuration to test locally the provider
terraform {
    required_providers {
        netbox = {
            source  = "terraform.example.com/local/netbox"
        }
    }
}