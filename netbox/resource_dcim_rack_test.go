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

func TestAccDcimRack_basic(t *testing.T) {
	name := "test rack"
	site_id := "12"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcimRackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDcimRackConfigBasic(name, site_id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcimRackExists("netbox_dcim_rack.test"),
				),
			},
		},
	})
}

func testAccCheckDcimRackDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_rack" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimRacksReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Dcim.DcimRacksRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Rack ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDcimRackExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rack ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimRacksReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Dcim.DcimRacksRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDcimRackConfigBasic(name string, site_id string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "Test"
  slug = "test"
}

resource "netbox_dcim_rack" "test" {
	name = "%s"
	site_id = "%s"
	facility = "tf facility rack"
	tenant_id = 6
	status = "available" 
	role_id = 2
	serial = "test serial" 
	asset_tag = netbox_tag.test.name
	type = "2-post-frame"
	width = 23
	u_height = 15
	desc_units = true
	outer_width = 11
	outer_depth = 10
	outer_unit = "mm"
	comments = "new comment"
	
	tags {
	  name = netbox_tag.test.name
	  slug = netbox_tag.test.slug
	}

	custom_fields = {
		rackCustomField = "rackCustomeFieldValue"
	}
  }
`, name, site_id)
}
