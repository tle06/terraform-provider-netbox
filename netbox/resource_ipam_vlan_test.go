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

func TestAccIpamVlan_basic(t *testing.T) {
	vid := "300"
	name := "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamVlanConfigBasic(name, vid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamVlanExists("netbox_ipam_vlan.test"),
				),
			},
		},
	})
}

func testAccCheckIpamVlanDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_vlan" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamVlansReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Ipam.IpamVlansRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Vlan ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamVlanExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vlan ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamVlansReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Ipam.IpamVlansRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamVlanConfigBasic(name string, vid string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_vlan" "test" {
	name = "%s"
	vid = "%s"

	}
`, name, vid)
}
