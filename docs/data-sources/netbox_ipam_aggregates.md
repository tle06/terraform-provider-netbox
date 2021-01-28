# netbox_ipam_aggregates Data Source

Use this data source to get information about all aggregates.

## Example Usage

```hcl
data "netbox_ipam_aggregates" "example" {
  family = 6
}
```

## Argument Reference

* `prefix` - (Optional) - A prefix used to filter the results.

* `family` - (Optional) - A value for the address family. Possible values include: `4` and `6`.

## Attribute Reference

* `results` - One or more `results` blocks as defined below.

The `results` block contains:

* `id` - The ID of the prefix.

* `prefix` - The address prefix.

* `description` - A description for the prefix.

* `family` - A `family` block as defined below.

The `family` block contains:

* `value` - A value for the address family. Possible values are: `4` and `6`.

* `label` - A label for the address family. Possible values are: `IPv4` and `IPv6`.
