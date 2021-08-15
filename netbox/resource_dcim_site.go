package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimSiteCreate,
		ReadContext:   resourceDcimSiteRead,
		UpdateContext: resourceDcimSiteUpdate,
		DeleteContext: resourceDcimSiteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringLenBetween(0, 100),
			},

			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.SiteStatusValueActive,
					models.SiteStatusValueDecommissioning,
					models.SiteStatusValuePlanned,
					models.SiteStatusValueRetired,
					models.SiteStatusValueStaging,
				}),
			},

			"region_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"facility": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"asn_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},
			"physical_address": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},
			"shipping_address": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},
			"latitude": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(9, 9),
			},
			"longitude": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(9, 9),
			},
			"contact_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 50),
			},
			"contact_phone": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 20),
			},
			"contact_email": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 254),
			},
			"comments": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
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

func resourceDcimSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	var name = d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimSitesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("region_id"); ok {
		regionID := int64(v.(int))
		params.Data.Region = &regionID
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("facility"); ok {
		params.Data.Facility = v.(string)
	}

	if v, ok := d.GetOk("asn_id"); ok {
		asnID := int64(v.(int))
		params.Data.Asn = &asnID
	}

	if v, ok := d.GetOk("time_zone"); ok {
		params.Data.TimeZone = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("physical_address"); ok {
		params.Data.PhysicalAddress = v.(string)
	}

	if v, ok := d.GetOk("shipping_address"); ok {
		params.Data.ShippingAddress = v.(string)
	}

	if v, ok := d.GetOk("latitude"); ok {
		latitude := v.(string)
		params.Data.Latitude = &latitude
	}

	if v, ok := d.GetOk("longitude"); ok {
		longitude := v.(string)
		params.Data.Longitude = &longitude
	}

	if v, ok := d.GetOk("contact_name"); ok {
		params.Data.ContactName = v.(string)
	}

	if v, ok := d.GetOk("contact_phone"); ok {
		params.Data.ContactPhone = v.(string)
	}

	if v, ok := d.GetOk("contact_email"); ok {
		params.Data.ContactEmail = v.(strfmt.Email)
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create site: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceDcimSiteRead(ctx, d, m)

	return diags
}

func resourceDcimSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimSitesReadParams{
		Context: ctx,
		ID:      siteID,
	}

	resp, err := c.Dcim.DcimSitesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get site: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("slug", resp.Payload.Slug)

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Region != nil {
		d.Set("region_id", resp.Payload.Region.ID)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Facility != "" {
		d.Set("facility", resp.Payload.Facility)
	}

	if resp.Payload.Asn != nil {
		d.Set("asn_id", resp.Payload.Asn)
	}

	if resp.Payload.TimeZone != "" {
		d.Set("time_zone", resp.Payload.TimeZone)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.PhysicalAddress != "" {
		d.Set("physical_address", resp.Payload.PhysicalAddress)
	}

	if resp.Payload.ShippingAddress != "" {
		d.Set("shipping_address", resp.Payload.ShippingAddress)
	}

	if resp.Payload.Latitude != nil {
		d.Set("latitude", resp.Payload.Latitude)
	}

	if resp.Payload.Longitude != nil {
		d.Set("longitude", resp.Payload.Longitude)
	}

	if resp.Payload.ContactName != "" {
		d.Set("contact_name", resp.Payload.ContactName)
	}

	if resp.Payload.ContactPhone != "" {
		d.Set("contact_phone", resp.Payload.ContactPhone)
	}

	if resp.Payload.ContactEmail != "" {
		d.Set("contact_email", resp.Payload.ContactEmail)
	}

	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceDcimSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimSitesPartialUpdateParams{
		Context: ctx,
		ID:      siteID,
	}

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("region_id") {
		regionID := int64(d.Get("region_id").(int))
		params.Data.Region = &regionID
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("facility") {
		params.Data.Facility = d.Get("facility").(string)
	}

	if d.HasChange("asn_id") {
		asnID := int64(d.Get("asn_id").(int))
		params.Data.Asn = &asnID
	}

	if d.HasChange("time_zone") {
		params.Data.TimeZone = d.Get("time_zone").(string)
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("physical_address") {
		params.Data.PhysicalAddress = d.Get("physical_address").(string)
	}

	if d.HasChange("shipping_address") {
		params.Data.ShippingAddress = d.Get("shipping_address").(string)
	}

	if d.HasChange("latitude") {
		latitude := d.Get("latitude").(string)
		params.Data.Latitude = &latitude
	}

	if d.HasChange("longitude") {
		longitude := d.Get("longitude").(string)
		params.Data.Longitude = &longitude
	}

	if d.HasChange("contact_name") {
		params.Data.ContactName = d.Get("contact_name").(string)
	}

	if d.HasChange("contact_phone") {
		params.Data.ContactPhone = d.Get("contact_phone").(string)
	}

	if d.HasChange("contact_email") {
		params.Data.ContactEmail = d.Get("contact_email").(strfmt.Email)
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

	_, err = c.Dcim.DcimSitesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update site: %v", err)
	}

	return resourceIpamPrefixRead(ctx, d, m)
}

func resourceDcimSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimSitesDeleteParams{
		Context: ctx,
		ID:      siteID,
	}

	_, err = c.Dcim.DcimSitesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete site: %v", err)
	}

	d.SetId("")

	return diags
}
