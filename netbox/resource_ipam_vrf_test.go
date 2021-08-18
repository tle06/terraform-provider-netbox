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

func TestAccIpamVRF_basic(t *testing.T) {
	name := "Test VRF"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamVRFDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamVRFConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamVRFExists("netbox_ipam_vrf.test"),
				),
			},
		},
	})
}

func testAccCheckIpamVRFDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_vrf" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamVrfsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Ipam.IpamVrfsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("VRF ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamVRFExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VRF ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamVrfsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Ipam.IpamVrfsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamVRFConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_vrf" "test" {
	name = "%s"
	}
`, name)
}
