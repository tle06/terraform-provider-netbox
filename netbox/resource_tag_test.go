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
	"github.com/netbox-community/go-netbox/netbox/client/extras"
)

func TestAccTag_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckTagConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists("netbox_tag.test"),
				),
			},
		},
	})
}

func testAccCheckTagDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_tag" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &extras.ExtrasTagsReadParams{
			Context: context.Background(),
			ID:      id,
		}

		resp, err := c.Extras.ExtrasTagsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Tag still exists (ID: %d)", resp.Payload.ID)
	}

	return nil
}

func testAccCheckTagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tag set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &extras.ExtrasTagsReadParams{
			Context: context.Background(),
			ID:      id,
		}

		_, err = c.Extras.ExtrasTagsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckTagConfigBasic() string {
	return `
resource "netbox_tag" "test" {
  name        = "Test"
  slug        = "test"
  color       = "ff0000"
  description = "Acceptance test"
}
`
}
