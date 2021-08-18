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
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func TestAccDcimRegion_basic(t *testing.T) {
	name := "test region"
	slug := "test-region"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcimRegionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDcimRegionConfigBasic(name, slug),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcimRegionExists("netbox_dcim_region.test"),
				),
			},
		},
	})
}

func testAccCheckDcimRegionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_region" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimRegionsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Dcim.DcimRegionsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Region ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDcimRegionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No region ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimRegionsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Dcim.DcimRegionsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDcimRegionConfigBasic(name string, slug string) string {
	return fmt.Sprintf(`
	resource "netbox_dcim_region" "example" {
		name = "%s"
		slug = "%s"
	  }
`, name, slug)
}
