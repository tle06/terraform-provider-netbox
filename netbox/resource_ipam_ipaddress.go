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

func resourceIpamIPAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamIPAddressCreate,
		ReadContext:   resourceIpamIPAddressRead,
		UpdateContext: resourceIpamIPAddressUpdate,
		DeleteContext: resourceIpamIPAddressDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},

			"nat_outside_id": {
				Type:     schema.TypeInt,
				Optional: true,
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

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.IPAddressStatusValueActive,
					models.IPAddressStatusValueDeprecated,
					models.IPAddressStatusValueDhcp,
					models.IPAddressStatusValueReserved,
					models.IPAddressStatusValueSlaac,
				}),

				Default: models.IPAddressStatusValueActive,
			},

			"role": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.IPAddressRoleValueAnycast,
					models.IPAddressRoleValueCarp,
					models.IPAddressRoleValueGlbp,
					models.IPAddressRoleValueHsrp,
					models.IPAddressRoleValueLoopback,
					models.IPAddressRoleValueSecondary,
					models.IPAddressRoleValueVip,
					models.IPAddressRoleValueVrrp,
				}),
			},

			"assigned_object_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"assigned_object_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vrf_id": {
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

func resourceIpamIPAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	address := d.Get("address").(string)
	nat_outside_id := int64(d.Get("nat_outside_id").(int))

	params := &ipam.IpamIPAddressesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableIPAddress{
		Address:    &address,
		NatOutside: &nat_outside_id,
		Tags:       expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("role"); ok {
		params.Data.Role = v.(string)
	}

	if v, ok := d.GetOk("vrf_id"); ok {
		vrfID := int64(v.(int))
		params.Data.Vrf = &vrfID
	}

	if v, ok := d.GetOk("assigned_object_id"); ok {
		assignedObjectID := int64(v.(int))
		params.Data.AssignedObjectID = &assignedObjectID
	}

	if v, ok := d.GetOk("assigned_object_type"); ok {
		params.Data.AssignedObjectType = v.(string)
	}

	if v, ok := d.GetOk("dns_name"); ok {
		params.Data.DNSName = v.(string)
	}

	if v, ok := d.GetOk("nat_inside_id"); ok {
		natInside := int64(v.(int))
		params.Data.NatInside = &natInside
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Ipam.IpamIPAddressesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create address: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamIPAddressRead(ctx, d, m)

	return diags
}

func resourceIpamIPAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamIPAddressesReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get address: %v", err)
	}

	d.Set("address", resp.Payload.Address)

	if resp.Payload.NatOutside != nil {
		d.Set("nat_outside_id", resp.Payload.NatOutside.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Role != nil {
		d.Set("role", resp.Payload.Role)
	}

	if resp.Payload.Vrf != nil {
		d.Set("vrf_id", resp.Payload.Vrf.ID)
	}

	if resp.Payload.AssignedObjectID != nil {
		d.Set("assigned_object_id", resp.Payload.AssignedObjectID)
	}
	if resp.Payload.AssignedObjectType != "" {
		d.Set("assigned_object_type", resp.Payload.AssignedObjectType)
	}
	if resp.Payload.DNSName != "" {
		d.Set("dns_name", resp.Payload.DNSName)
	}
	if resp.Payload.NatInside != nil {
		d.Set("nat_inside_id", resp.Payload.NatInside)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceIpamIPAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	address := d.Get("address").(string)

	params := &ipam.IpamIPAddressesPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.WritableIPAddress{
		Address: &address,
	}

	if d.HasChange("nat_outside_id") {
		natOutside := int64(d.Get("nat_outside_id").(int))
		params.Data.NatOutside = &natOutside
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Vrf = &tenantID
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("role") {
		params.Data.Role = d.Get("role").(string)
	}

	if d.HasChange("vrf_id") {
		vrfID := int64(d.Get("vrf_id").(int))
		params.Data.Vrf = &vrfID
	}

	if d.HasChange("assigned_object_id") {
		assignedObjectID := int64(d.Get("assigned_object_id").(int))
		params.Data.AssignedObjectID = &assignedObjectID
	}

	if d.HasChange("assigned_object_type") {
		params.Data.AssignedObjectType = d.Get("assigned_object_type").(string)
	}

	if d.HasChange("dns_name") {
		params.Data.DNSName = d.Get("dns_name").(string)
	}

	if d.HasChange("nat_inside_id") {
		natInside := int64(d.Get("nat_inside_id").(int))
		params.Data.NatInside = &natInside
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	_, err = c.Ipam.IpamIPAddressesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update address: %v", err)
	}

	return resourceIpamIPAddressRead(ctx, d, m)
}

func resourceIpamIPAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamIPAddressesDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Ipam.IpamIPAddressesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete address: %v", err)
	}

	d.SetId("")

	return diags
}
