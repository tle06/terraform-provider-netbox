package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/tenancy"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTenancyTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTenancyTenantCreate,
		ReadContext:   resourceTenancyTenantRead,
		UpdateContext: resourceTenancyTenantUpdate,
		DeleteContext: resourceTenancyTenantDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},

			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"group_id": {
				Type:     schema.TypeInt,
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

func resourceTenancyTenantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	var name = d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &tenancy.TenancyTenantsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableTenant{
		Name: &name,
		Slug: &slug,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("group_id"); ok {
		groupID := int64(v.(int))
		params.Data.Group = &groupID
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	resp, err := c.Tenancy.TenancyTenantsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create tenant: %v", err)
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceTenancyTenantRead(ctx, d, m)

	return diags
}

func resourceTenancyTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &tenancy.TenancyTenantsReadParams{
		Context: ctx,
		ID:      tenantID,
	}

	resp, err := c.Tenancy.TenancyTenantsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get tenant: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("slug", resp.Payload.Slug)

	if resp.Payload.Group != nil {
		d.Set("group_id", resp.Payload.Group.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceTenancyTenantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &tenancy.TenancyTenantsPartialUpdateParams{
		Context: ctx,
		ID:      tenantID,
	}

	params.Data = &models.WritableTenant{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("group_id") {
		groupID := int64(d.Get("group_id").(int))
		params.Data.Group = &groupID
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("comments") {
		params.Data.Comments = d.Get("comments").(string)
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	_, err = c.Tenancy.TenancyTenantsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update tenant: %v", err)
	}

	return resourceTenancyTenantRead(ctx, d, m)
}

func resourceTenancyTenantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	regionID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &tenancy.TenancyTenantsDeleteParams{
		Context: ctx,
		ID:      regionID,
	}

	_, err = c.Tenancy.TenancyTenantsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete tenant: %v", err)
	}

	d.SetId("")

	return diags
}
