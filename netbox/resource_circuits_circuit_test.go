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
	"github.com/netbox-community/go-netbox/netbox/client/circuits"
)

func TestAccCircuitsCircuit_basic(t *testing.T) {
	cid := "test circuit"
	provider_id := "2"
	type_id := "1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCircuitsCircuitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCircuitsCircuitConfigBasic(cid, type_id, provider_id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCircuitsCircuitExists("netbox_circuits_circuit.test"),
				),
			},
		},
	})
}

func testAccCheckCircuitsCircuitDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_circuits_circuit" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &circuits.CircuitsCircuitsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Circuits.CircuitsCircuitsRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Circuit ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckCircuitsCircuitExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No circuit ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &circuits.CircuitsCircuitsReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Circuits.CircuitsCircuitsRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckCircuitsCircuitConfigBasic(cid string, type_id string, provider_id string) string {
	return fmt.Sprintf(`
resource "netbox_circuits_circuit" "test" {

	cid ="%s"
	type_id = "%s"
	provider_id = "%s"
}

`, cid, type_id, provider_id)
}
