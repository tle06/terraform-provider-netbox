package netbox

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func dataSourceIpamAvailablePrefixes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamAvailablePrefixesRead,
		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"family": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"vrf": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"rd": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIpamAvailablePrefixesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	var diags diag.Diagnostics

	params := &ipam.IpamPrefixesAvailablePrefixesReadParams{
		Context: ctx,
		ID:      int64(d.Get("prefix_id").(int)),
	}

	resp, err := c.Ipam.IpamPrefixesAvailablePrefixesRead(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get available prefixes: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("prefixes", flattenIpamAvailablePrefixes(resp.Payload))

	return diags
}

func flattenIpamAvailablePrefixes(input []*models.AvailablePrefix) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["family"] = item.Family
		values["prefix"] = item.Prefix
		values["vrf"] = flattenIpamPrefixVRF(item.Vrf)

		result = append(result, values)
	}

	return result
}
