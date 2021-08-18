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
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
)

func TestAccIpamRir_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamRirDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamRirConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamRirExists("netbox_ipam_rir.test"),
				),
			},
		},
	})
}

func testAccCheckIpamRirDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_rir" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamRirsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Ipam.IpamRirsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("RIR ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamRirExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RIR ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamRirsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Ipam.IpamRirsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamRirConfigBasic() string {
	return `
resource "netbox_ipam_rir" "test" {
  name        = "Test"
  slug        = "test"
}
`
}
