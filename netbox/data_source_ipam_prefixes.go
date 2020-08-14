package netbox

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/innovationnorway/go-netbox/models"
	"github.com/innovationnorway/go-netbox/plumbing"
	"github.com/innovationnorway/go-netbox/plumbing/ipam"
)

func dataSourceIpamPrefixes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamPrefixesRead,

		Schema: map[string]*schema.Schema{
			"contains": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mask_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"site": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tag": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"within": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"within_include": {
				Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func dataSourceIpamPrefixesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*plumbing.Netbox)

	var diags diag.Diagnostics

	params := &ipam.IpamPrefixesListParams{
		Context: ctx,
	}

	if v, ok := d.GetOk("contains"); ok {
		contains := v.(string)
		params.Contains = &contains
	}

	if v, ok := d.GetOk("mask_length"); ok {
		maskLength := float64(v.(int))
		params.MaskLength = &maskLength
	}

	if v, ok := d.GetOk("prefix"); ok {
		prefix := v.(string)
		params.Prefix = &prefix
	}

	if v, ok := d.GetOk("region"); ok {
		region := v.(string)
		params.Region = &region
	}

	if v, ok := d.GetOk("region"); ok {
		region := v.(string)
		params.Region = &region
	}

	if v, ok := d.GetOk("role"); ok {
		role := v.(string)
		params.Role = &role
	}

	if v, ok := d.GetOk("site"); ok {
		site := v.(string)
		params.Site = &site
	}

	if v, ok := d.GetOk("status"); ok {
		status := v.(string)
		params.Status = &status
	}

	if v, ok := d.GetOk("tag"); ok {
		tag := v.(string)
		params.Tag = &tag
	}

	if v, ok := d.GetOk("tenant"); ok {
		tenant := v.(string)
		params.Tenant = &tenant
	}

	if v, ok := d.GetOk("within"); ok {
		within := v.(string)
		params.Within = &within
	}

	if v, ok := d.GetOk("within_include"); ok {
		withinInclude := v.(string)
		params.WithinInclude = &withinInclude
	}

	resp, err := c.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get prefixes: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenIpamPrefixesResults(resp.Payload.Results))

	return diags
}

func flattenIpamPrefixesResults(input []*models.Prefix) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["family"] = flattenIpamPrefixFamily(item.Family)
		values["prefix"] = item.Prefix
		values["site"] = flattenIpamPrefixSite(item.Site)
		values["vrf"] = flattenIpamPrefixVRF(item.Vrf)
		values["tenant"] = flattenIpamPrefixTenant(item.Tenant)
		values["vlan"] = flattenIpamPrefixVLAN(item.Vlan)
		values["status"] = flattenIpamPrefixStatus(item.Status)
		values["role"] = flattenIpamPrefixRole(item.Role)
		values["is_pool"] = item.IsPool
		values["description"] = item.Description
		values["tags"] = item.Tags
		values["custom_fields"] = item.CustomFields

		result = append(result, values)
	}

	return result
}
