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

func TestAccDcimInterface_basic(t *testing.T) {
	name := "test interface"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcimInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDcimInterfaceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcimInterfaceExists("netbox_dcim_interface.test"),
				),
			},
		},
	})
}

func testAccCheckDcimInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_interface" {
			continue
		}

		prefixID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimInterfacesReadParams{
			Context: context.Background(),
			ID:      prefixID,
		}

		resp, err := c.Dcim.DcimInterfacesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Interface ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDcimInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		interfaceID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimInterfacesReadParams{
			Context: context.Background(),
			ID:      interfaceID,
		}

		_, err = c.Dcim.DcimInterfacesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDcimInterfaceConfigBasic(name string) string {
	return fmt.Sprintf(`

resource "netbox_tag" "test-interface" {
	name = "Test Interface"
	slug = "test-interface"
  }
	
resource "netbox_dcim_site" "test-interface" {
	name = "test-interface"
	slug = "test-interface"
	status = "active"
}

resource "netbox_dcim_rack" "test-interface" {
	name = "rack-test-interface"
	site_id = netbox_dcim_site.test-interface.id

}

resource "netbox_dcim_device" "test-interface" {
	device_type_id = 7
	device_role_id = 4
	site_id = netbox_dcim_site.test-interface.id

}

resource "netbox_dcim_interface" "test" {

	device_id = netbox_dcim_device.test-interface.id
	type = "virtual"
	name = "%s"
	tagged_vlan = [64]
}

`, name)
}
