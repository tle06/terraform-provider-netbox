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

func TestAccIpamAggregate_basic(t *testing.T) {
	prefix := "100.0.0.0/16"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpamAggregateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIpamAggregateConfigBasic(prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamAggregateExists("netbox_ipam_aggregates.test"),
				),
			},
		},
	})
}

func testAccCheckIpamAggregateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_ipam_aggregate" {
			continue
		}

		aggregateID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamAggregatesReadParams{
			Context: context.Background(),
			ID:      aggregateID,
		}

		resp, err := c.Ipam.IpamAggregatesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Aggregate ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckIpamAggregateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Aggregate ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		aggregateID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &ipam.IpamAggregatesReadParams{
			Context: context.Background(),
			ID:      aggregateID,
		}

		_, err = c.Ipam.IpamAggregatesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIpamAggregateConfigBasic(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_rir" "test" {
	name        = "Test"
	slug        = "test"
	}
resource "netbox_ipam_aggregates" "test" {
  prefix      = "%s"
  rir_id         = netbox_ipam_rir.test.id

}
`, prefix)
}
