package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamPrefix_basic(t *testing.T) {
	prefix := "10.0.0.0/16"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamPrefixConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_prefix.test", "prefix", prefix),
				),
			},
		},
	})
}

func testAccDataSourceIpamPrefixConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "test" {
  prefix = "%s"
  status = "active"
}

data "netbox_ipam_prefix" "test" {
  prefix_id = netbox_ipam_prefix.test.id
}
`, prefix)
}
