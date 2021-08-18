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

func TestAccDcimSite_basic(t *testing.T) {
	name := "test site"
	slug := "test-site"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcimSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDcimSiteConfigBasic(name, slug),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcimSiteExists("netbox_dcim_site.test"),
				),
			},
		},
	})
}

func testAccCheckDcimSiteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_site" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimSitesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Dcim.DcimSitesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Site ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDcimSiteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No site ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimSitesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Dcim.DcimSitesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDcimSiteConfigBasic(name string, slug string) string {
	return fmt.Sprintf(`
resource "netbox_extras_tag" "test" {
  name = "Test"
  slug = "test"
}

resource "netbox_dcim_site" "test" {
	name = "%s"
	slug = "%s"
	status = "planned"
	facility = "tf facility"
	asn_id = 65000
	time_zone = "Africa/Asmara"
	description = "Acceptance test"
	physical_address = "my physicla address"
	shipping_address = "my shipping address"
	latitude = "10.800000"
	longitude = "11.600000"
	contact_name = "John doe"
	contact_phone = "+33 7 45 81 81 93"
	comments = "this is a comment"
	custom_fields = {
		tf-test = "customField"
	  }
	tags  {
	  name = netbox_extras_tag.test.name
	  slug = netbox_extras_tag.test.slug
	}
  }
`, name, slug)
}
