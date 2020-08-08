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

func resourceIpamPrefix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamPrefixCreate,
		ReadContext:   resourceIpamPrefixRead,
		UpdateContext: resourceIpamPrefixUpdate,
		DeleteContext: resourceIpamPrefixDelete,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isCIDR,
			},

			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},

			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.PrefixStatusValueActive,
					models.PrefixStatusValueContainer,
					models.PrefixStatusValueDeprecated,
					models.PrefixStatusValueReserved,
				}),
			},

			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"family": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIpamPrefixCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	var diags diag.Diagnostics

	prefix := d.Get("prefix").(string)

	params := &ipam.IpamPrefixesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritablePrefix{
		Prefix: &prefix,
		Tags:   expandStringSlice(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("site_id"); ok {
		siteID := int64(v.(int))
		params.Data.Site = &siteID
	}

	if v, ok := d.GetOk("vrf_id"); ok {
		vrfID := int64(v.(int))
		params.Data.Vrf = &vrfID
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("vlan_id"); ok {
		vlanID := int64(v.(int))
		params.Data.Vlan = &vlanID
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("role_id"); ok {
		roleID := int64(v.(int))
		params.Data.Role = &roleID
	}

	if v, ok := d.GetOk("is_pool"); ok {
		params.Data.IsPool = v.(bool)
	}

	resp, err := c.Ipam.IpamPrefixesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create prefix: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamPrefixRead(ctx, d, m)

	return diags
}

func resourceIpamPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamPrefixesReadParams{
		Context: ctx,
		ID:      prefixID,
	}

	resp, err := c.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get prefix: %v", err)
	}

	d.Set("family", resp.Payload.Family.Label)
	d.Set("prefix", resp.Payload.Prefix)
	d.Set("description", resp.Payload.Description)

	if resp.Payload.Site != nil {
		d.Set("site_id", resp.Payload.Site.ID)
	}

	if resp.Payload.Vrf != nil {
		d.Set("vrf_id", resp.Payload.Vrf.ID)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Vlan != nil {
		d.Set("vlan_id", resp.Payload.Vlan.ID)
	}

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Role != nil {
		d.Set("role_id", resp.Payload.Role.ID)
	}

	d.Set("is_pool", resp.Payload.IsPool)
	d.Set("tags", resp.Payload.Tags)
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceIpamPrefixUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	prefix := d.Get("prefix").(string)

	params := &ipam.IpamPrefixesPartialUpdateParams{
		Context: ctx,
		ID:      prefixID,
	}

	params.Data = &models.WritablePrefix{
		Prefix: &prefix,
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		params.Data.Site = &siteID
	}

	if d.HasChange("vrf_id") {
		vrfID := int64(d.Get("vrf_id").(int))
		params.Data.Vrf = &vrfID
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Vrf = &tenantID
	}

	if d.HasChange("vlan_id") {
		vlanID := int64(d.Get("vlan_id").(int))
		params.Data.Vlan = &vlanID
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("role_id") {
		roleID := int64(d.Get("role_id").(int))
		params.Data.Role = &roleID
	}

	if d.HasChange("is_pool") {
		params.Data.IsPool = d.Get("is_pool").(bool)
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandStringSlice(d.Get("tags").([]interface{}))
	}

	_, err = c.Ipam.IpamPrefixesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update prefix: %v", err)
	}

	return resourceIpamPrefixRead(ctx, d, m)
}

func resourceIpamPrefixDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBox)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamPrefixesDeleteParams{
		Context: ctx,
		ID:      prefixID,
	}

	_, err = c.Ipam.IpamPrefixesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete prefix: %v", err)
	}

	d.SetId("")

	return diags
}
