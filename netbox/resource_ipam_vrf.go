package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpamVRF() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamVRFCreate,
		ReadContext:   resourceIpamVRFRead,
		UpdateContext: resourceIpamVRFUpdate,
		DeleteContext: resourceIpamVRFDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"enforce_unique": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"rd": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIpamVRFCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	params := &ipam.IpamVrfsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableVRF{
		Name: &name,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("enforce_unique"); ok {
		params.Data.EnforceUnique = v.(bool)
	}

	if v, ok := d.GetOk("rd"); ok {
		rd := v.(string)
		params.Data.Rd = &rd
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Ipam.IpamVrfsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create vrf: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamVRFRead(ctx, d, m)

	return diags
}

func resourceIpamVRFRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVrfsReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Ipam.IpamVrfsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get vrf: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("enforce_unique", resp.Payload.EnforceUnique)

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.Rd != nil {
		d.Set("rd", resp.Payload.Rd)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceIpamVRFUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)

	params := &ipam.IpamVrfsPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.WritableVRF{
		Name: &name,
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("enforce_unique") {
		params.Data.EnforceUnique = d.Get("enforce_unique").(bool)
	}

	if d.HasChange("rd") {
		rd := d.Get("rd").(string)
		params.Data.Rd = &rd
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	_, err = c.Ipam.IpamVrfsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update vrf: %v", err)
	}

	return resourceIpamVRFRead(ctx, d, m)
}

func resourceIpamVRFDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVrfsDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Ipam.IpamVrfsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete vrf: %v", err)
	}

	d.SetId("")

	return diags
}
