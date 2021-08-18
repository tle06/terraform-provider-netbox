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

func TestAccIpamPrefix_basic(t *testing.T) {
	prefix := "10.0.0.0/16"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamPrefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamPrefixConfigBasic(prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamPrefixExists("netbox_ipam_prefix.test"),
				),
			},
		},
	})
}

func testAccCheckIpamPrefixDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_prefix" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamPrefixesReadParams{
			Context: context.Background(),
			ID:      objectID,
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

func testAccCheckIpamPrefixExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Prefix ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamPrefixesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Ipam.IpamPrefixesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamPrefixConfigBasic(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_extras_tag" "test" {
  name = "Test"
  slug = "test"
}

resource "netbox_ipam_prefix" "test" {
  prefix      = "%s"
  description = "Acceptance test"
  is_pool     = false
  tags {
	name = netbox_extras_tag.test.name
    slug = netbox_extras_tag.test.slug
  }
  status = "active"
}
`, prefix)
}
