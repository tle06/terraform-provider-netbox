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

func dataSourceIpamAggregates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamAggregatesRead,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"family": {
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"family": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

					},
				},
			},
		},
	}
}

func dataSourceIpamAggregatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &ipam.IpamAggregatesListParams{
		Context: ctx,
	}

	if v, ok := d.GetOk("prefix"); ok {
		prefix := v.(string)
		params.Prefix = &prefix
	}

	if v, ok := d.GetOk("family"); ok {
		family := v.(float64)
		params.Family = &family
	}

	resp, err := c.Ipam.IpamAggregatesList(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get prefixes: %v", err)
	}

	//lintignore:R017
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenIpamAggregatesResults(resp.Payload.Results))

	return diags
}

func flattenIpamAggregatesResults(input []*models.Aggregate) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["family"] = flattenIpamAggregateFamily(item.Family)
		values["prefix"] = item.Prefix
		values["description"] = item.Description
		
		result = append(result, values)
	}

	return result
}

func flattenIpamAggregateFamily(input *models.AggregateFamily) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["label"] = input.Label
	values["value"] = input.Value

	return []interface{}{values}
}