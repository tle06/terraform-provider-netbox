package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamPrefixes_basic(t *testing.T) {
	prefix := "10.0.0.0/16"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamPrefixesConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_prefixes.test", "results.0.prefix", prefix),
				),
			},
		},
	})
}

func testAccDataSourceIpamPrefixesConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "test" {
  prefix = "%s"
  status = "active"
}

data "netbox_ipam_prefixes" "test" {
  prefix = netbox_ipam_prefix.test.prefix
  depends_on = [
    netbox_ipam_prefix.test,
  ]
}
`, prefix)
}
