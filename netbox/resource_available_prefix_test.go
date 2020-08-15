package netbox

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/innovationnorway/go-netbox/plumbing"
	"github.com/innovationnorway/go-netbox/plumbing/ipam"
)

func TestAccIpamAvailablePrefix_basic(t *testing.T) {
	prefix := "10.0.0.0/16"
	length := 24

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamAvailablePrefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamAvailablePrefixConfigBasic(prefix, length),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamAvailablePrefixExists("netbox_ipam_available_prefix.test"),
				),
			},
		},
	})
}

func testAccCheckIpamAvailablePrefixDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*plumbing.Netbox)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_available_prefix" {
			continue
		}

		prefixID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamPrefixesReadParams{
			Context: context.Background(),
			ID:      prefixID,
		}

		resp, err := c.Ipam.IpamPrefixesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Prefix ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamAvailablePrefixExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Prefix ID set")
		}

		c := testAccProvider.Meta().(*plumbing.Netbox)

		prefixID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamPrefixesReadParams{
			Context: context.Background(),
			ID:      prefixID,
		}

		_, err = c.Ipam.IpamPrefixesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamAvailablePrefixConfigBasic(prefix string, length int) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "test" {
  prefix = "%s"
  status = "active"
}

resource "netbox_ipam_available_prefix" "test" {
  prefix_id     = netbox_ipam_prefix.test.id
  prefix_length = "%d"
}
`, prefix, length)
}
