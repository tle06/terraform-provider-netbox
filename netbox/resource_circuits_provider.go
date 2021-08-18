package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/circuits"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCircuitsProvider() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCircuitsProviderCreate,
		ReadContext:   resourceCircuitsProviderRead,
		UpdateContext: resourceCircuitsProviderUpdate,
		DeleteContext: resourceCircuitsProviderDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"asn": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"account": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"admin_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"noc_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"portal_url": {
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

func resourceCircuitsProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &circuits.CircuitsProvidersCreateParams{
		Context: ctx,
	}

	params.Data = &models.Provider{
		Name: &name,
		Slug: &slug,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("account"); ok {
		params.Data.Account = v.(string)
	}

	if v, ok := d.GetOk("asn"); ok {
		asn := int64(v.(int))
		params.Data.Asn = &asn
	}

	if v, ok := d.GetOk("admin_contact"); ok {
		params.Data.AdminContact = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("noc_contact"); ok {
		params.Data.NocContact = v.(string)
	}

	if v, ok := d.GetOk("portal_url"); ok {
		portalURL := strfmt.URI(v.(string))
		params.Data.PortalURL = portalURL
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Circuits.CircuitsProvidersCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create circuit: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceCircuitsProviderRead(ctx, d, m)

	return diags
}

func resourceCircuitsProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &circuits.CircuitsProvidersReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Circuits.CircuitsProvidersRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get provider: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("slug", resp.Payload.Slug)

	if resp.Payload.Account != "" {
		d.Set("account", resp.Payload.Account)
	}

	if resp.Payload.Asn != nil {
		d.Set("asn", resp.Payload.Asn)
	}
	if resp.Payload.AdminContact != "" {
		d.Set("admin_contact", resp.Payload.AdminContact)
	}

	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}
	if resp.Payload.NocContact != "" {
		d.Set("noc_contact", resp.Payload.NocContact)
	}
	if resp.Payload.PortalURL != "" {
		d.Set("portal_url", resp.Payload.PortalURL)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceCircuitsProviderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &circuits.CircuitsProvidersPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.Provider{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("account") {
		params.Data.Account = d.Get("account").(string)
	}

	if d.HasChange("asn") {
		asn := int64(d.Get("asn").(int))
		params.Data.Asn = &asn
	}

	if d.HasChange("admin_contact") {
		params.Data.AdminContact = d.Get("admin_contact").(string)
	}

	if d.HasChange("comments") {
		params.Data.Comments = d.Get("comments").(string)
	}

	if d.HasChange("noc_contact") {
		params.Data.NocContact = d.Get("noc_contact").(string)
	}

	if d.HasChange("portal_url") {
		portalURL := strfmt.URI(d.Get("portal_url").(string))
		params.Data.PortalURL = portalURL
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	_, err = c.Circuits.CircuitsProvidersPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update provider: %v", err)
	}

	return resourceCircuitsProviderRead(ctx, d, m)
}

func resourceCircuitsProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &circuits.CircuitsProvidersDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Circuits.CircuitsProvidersDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete provider: %v", err)
	}

	d.SetId("")

	return diags
}
