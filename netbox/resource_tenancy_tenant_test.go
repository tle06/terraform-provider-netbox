package netbox

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/tenancy"
)

func TestAccTenancyTenant_basic(t *testing.T) {
	name := "test"
	slug := "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTenancyTenantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckTenancyTenantConfigBasic(name, slug),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTenancyTenantExists("netbox_tenancy_tenant.test"),
				),
			},
		},
	})
}

func testAccCheckTenancyTenantDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_tenancy_tenant" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &tenancy.TenancyTenantsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Tenancy.TenancyTenantsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Tenant ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckTenancyTenantExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tenant ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &tenancy.TenancyTenantsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Tenancy.TenancyTenantsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckTenancyTenantConfigBasic(name string, slug string) string {
	return fmt.Sprintf(`
resource "netbox_tenancy_tenant" "test" {
	name = "%s"
	slug = "%s"

	custom_fields = {
		cust_id = ""
	}
}
`, name, slug)
}
