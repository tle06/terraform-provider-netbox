package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
)

func dataSourceIpamPrefix() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamPrefixRead,
		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
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

			"site": {
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
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

			"tenant": {
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
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"vlan": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"role": {
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
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"is_pool": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIpamPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	var diags diag.Diagnostics

	params := &ipam.IpamPrefixesReadParams{
		Context: ctx,
		ID:      int64(d.Get("prefix_id").(int)),
	}

	resp, err := c.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get prefix: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))
	d.Set("family", flattenIpamPrefixFamily(resp.Payload.Family))
	d.Set("prefix", resp.Payload.Prefix)
	d.Set("site", flattenIpamPrefixSite(resp.Payload.Site))
	d.Set("vrf", flattenIpamPrefixVRF(resp.Payload.Vrf))
	d.Set("tenant", flattenIpamPrefixTenant(resp.Payload.Tenant))
	d.Set("vlan", flattenIpamPrefixVLAN(resp.Payload.Vlan))
	d.Set("status", flattenIpamPrefixStatus(resp.Payload.Status))
	d.Set("role", flattenIpamPrefixRole(resp.Payload.Role))
	d.Set("is_pool", resp.Payload.IsPool)
	d.Set("description", resp.Payload.Description)
	d.Set("tags", resp.Payload.Tags)
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func flattenIpamPrefixFamily(input *models.PrefixFamily) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["label"] = input.Label
	values["value"] = input.Value

	return []interface{}{values}
}

func flattenIpamPrefixSite(input *models.NestedSite) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["name"] = input.Name
	values["slug"] = input.Slug

	return []interface{}{values}
}

func flattenIpamPrefixVRF(input *models.NestedVRF) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["name"] = input.Name
	values["rd"] = input.Rd

	return []interface{}{values}
}

func flattenIpamPrefixTenant(input *models.NestedTenant) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["name"] = input.Name
	values["slug"] = input.Slug

	return []interface{}{values}
}

func flattenIpamPrefixVLAN(input *models.NestedVLAN) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["vid"] = input.Vid
	values["name"] = input.Name
	values["display_name"] = input.DisplayName

	return []interface{}{values}
}

func flattenIpamPrefixStatus(input *models.PrefixStatus) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["label"] = input.Label
	values["value"] = input.Value

	return []interface{}{values}
}

func flattenIpamPrefixRole(input *models.NestedRole) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["name"] = input.Name
	values["slug"] = input.Slug

	return []interface{}{values}
}
