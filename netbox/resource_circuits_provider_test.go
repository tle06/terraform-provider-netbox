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
	"github.com/netbox-community/go-netbox/netbox/client/circuits"
)

func TestAccCircuitsProvider_basic(t *testing.T) {
	name := "test provider"
	slug := "test-provider"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCircuitsProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCircuitsProviderConfigBasic(name, slug),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCircuitsProviderExists("netbox_circuits_provider.test"),
				),
			},
		},
	})
}

func testAccCheckCircuitsProviderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_circuits_provider" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &circuits.CircuitsProvidersReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Circuits.CircuitsProvidersRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("provider ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckCircuitsProviderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No provider ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &circuits.CircuitsProvidersReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Circuits.CircuitsProvidersRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckCircuitsProviderConfigBasic(name string, slug string) string {
	return fmt.Sprintf(`
resource "netbox_circuits_provider" "test" {
	name = "%s"
	slug = "%s"
	}

`, name, slug)
}
