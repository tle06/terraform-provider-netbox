package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamAvailablePrefixes_basic(t *testing.T) {
	prefix := "10.0.0.0/16"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamAvailablePrefixesConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_available_prefixes.test", "prefixes.0.prefix", prefix),
				),
			},
		},
	})
}

func testAccDataSourceIpamAvailablePrefixesConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "test" {
  prefix = "%s"
  status = "active"
}

data "netbox_ipam_available_prefixes" "test" {
  prefix_id = netbox_ipam_prefix.test.id
}
`, prefix)
}
