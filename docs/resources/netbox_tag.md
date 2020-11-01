# netbox_tag Resource

Creates a tag.

## Example Usage

```hcl
resource "netbox_tag" "example" {
  name  = "Example"
  slug  = "example"
  color = "ff0000"
}
```

## Argument Reference

* `name` - (Required) The name of the name.

* `slug` - (Required) The tag of the tag.

* `color` - (Optional) The color of the tag.

* `description` - (Optional) A description of the tag.

## Attribute Reference

* `id` - The ID of the tag.
