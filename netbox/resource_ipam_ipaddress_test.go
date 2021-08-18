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

func TestAccIpamIPAddress_basic(t *testing.T) {
	address := "10.0.0.1/24"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamIPAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamIPAddressConfigBasic(address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamIPAddressExists("netbox_ipam_ipaddress.test"),
				),
			},
		},
	})
}

func testAccCheckIpamIPAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_ipaddress" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamIPAddressesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Ipam.IpamIPAddressesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("IP address ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamIPAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IP address ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamIPAddressesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Ipam.IpamIPAddressesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamIPAddressConfigBasic(address string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_ipaddress" "test" {
	address = "%s"
}
`, address)
}
