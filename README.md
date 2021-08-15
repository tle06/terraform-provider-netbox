# Terraform Provider for Netbox

## Requirements

- [NetBox](https://netbox.readthedocs.io/) >= 2.9
- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.16

## Usage

```hcl
provider "netbox" {
  host  = "http://localhost:8000"
  token = "66a48ac409ec56b3f345eee3d10a42fa2fc1b8b9"
}

resource "netbox_ipam_prefix" "example" {
  prefix = "10.0.0.0/16"
  status = "reserved"
}
```

## Contributing

To build the provider:

```sh
$ go build
```

To test the provider:

```sh
$ go test -v ./...
```

To run all acceptance tests:

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ TF_ACC=1 go test -v ./...
```

To run a subset of acceptance tests:

```sh
$ TF_ACC=1 go test -v ./... -run=TestAccIpamPrefix
```

The following environment variables must be set to run acceptance tests:

- `NETBOX_HOST`
- `NETBOX_TOKEN`

## Test the provider localy

### Windows environement

Build the provider from the root of the repo

```powershell
go build
```

Create the pluging structure folder

```powershell
New-Item -ItemType directory -Path "%APPADATA%\terraform.d\plugins\terraform.example.com\local\netbox\1.0\windows_amd64"
```

Move the package to the plugin directory

```powershell
Move-Item -Path .\terraform-provider-netbox.exe -Destination "$env:APPDATA\terraform.d\plugins\terraform.example.com\local\netbox\1.0\windows_amd64" -Force
```

Browse to the test terarform folder

```powershell
cd .\test\terraform\
```

Remove `.terraform` folder and `.terraform.lock.hcl`

```powershell
rm .\.terraform\ -Confirm:$False -Recurse; rm .\.terraform.lock.hcl -Confirm:$False
```

Terraform init

```powershell
terraform.exe init
```

Terraform plan

```powershell
terraform.exe plan
```

Terraform apply

```powershell
terraform.exe apply -auto-approve
```
