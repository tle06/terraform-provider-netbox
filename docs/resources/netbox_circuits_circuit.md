# netbox_ipam_prefix Resource

Creates a circuit.

## Example Usage

```hcl
resource "netbox_circuits_circuit" "example" {

  cid ="example"
  type_id = 1
  provider_id = 2


}
```

## Argument Reference

* `cid` - (Required) The cid of the circuit.

* `type_id` - (Required) The type IDlink to the circuit.

* `provider_id` - (Required) The provider ID link to the circuit

* `status` - (Optional) The status of the circuit. Possible value: `planned`, `provisioning`, `active`, `offline`, `deprovisioning`, `decommissioned`.

* `tenant_id` - (Optional) The tenant ID to add.

* `commit_rate` - (Optional) The commit rate (bandwidth) of the circuit.

* `comments` - (Optional) The comments on the circuit.

* `description` - (Optional) The description of the circuit.

* `install_date` - (Optional) The installation date of the circuit.

* `tags` - (Optional) List of tags to assign to the circuit. Each tag need to be input in a tags block and refering a resources previously created.
  ```
    tags {
      name = netbox_extras_tag.example2.name
      slug = netbox_extras_tag.example2.slug
    }
  ```

* `custom_fields` - (Optional) A mapping of custom fields to assign to the circuit. The custom fields need to be created before usage.
  ```
  custom_fields = {
    myNewCustomField = "customFieldValue"
  }
  ```

## Attribute Reference

* `id` - The prefix ID.
