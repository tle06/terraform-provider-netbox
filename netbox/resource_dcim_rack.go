package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimRack() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimRackCreate,
		ReadContext:   resourceDcimRackRead,
		UpdateContext: resourceDcimRackUpdate,
		DeleteContext: resourceDcimRackDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringLenBetween(1, 50),
			},

			"facility": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 50),
			},

			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.RackStatusValueActive,
					models.RackStatusValueAvailable,
					models.RackStatusValueDeprecated,
					models.RackStatusValuePlanned,
					models.RackStatusValueReserved,
				}),
				Default: models.RackStatusValueActive,
			},

			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"serial": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 50),
			},
			"asset_tag": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 50),
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.RackTypeValueNr2PostFrame,
					models.RackTypeValueNr4PostCabinet,
					models.RackTypeValueNr4PostFrame,
					models.RackTypeValueWallCabinet,
					models.RackTypeValueWallFrame,
				}),
			},
			"width": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  19,
			},

			"u_height": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  42,
			},

			"desc_units": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"outer_width": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"outer_depth": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"outer_unit": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.RackOuterUnitValueIn,
					models.RackOuterUnitValueMm,
				}),
				Default: models.RackOuterUnitValueMm,
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

func resourceDcimRackCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	var name = d.Get("name").(string)

	params := &dcim.DcimRacksCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableRack{
		Name: &name,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("facility"); ok {
		facilityID := v.(string)
		params.Data.FacilityID = &facilityID
	}

	if v, ok := d.GetOk("site_id"); ok {
		siteID := int64(v.(int))
		params.Data.Site = &siteID
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("role_id"); ok {
		roleID := int64(v.(int))
		params.Data.Role = &roleID
	}

	if v, ok := d.GetOk("serial"); ok {
		params.Data.Serial = v.(string)
	}

	if v, ok := d.GetOk("asset_tag"); ok {
		assetTag := v.(string)
		params.Data.AssetTag = &assetTag
	}

	if v, ok := d.GetOk("type"); ok {
		params.Data.Type = v.(string)
	}

	if v, ok := d.GetOk("width"); ok {
		params.Data.Width = int64(v.(int))
	}

	if v, ok := d.GetOk("u_height"); ok {
		params.Data.UHeight = int64(v.(int))
	}

	if v, ok := d.GetOk("desc_units"); ok {
		params.Data.DescUnits = v.(bool)
	}

	if v, ok := d.GetOk("outer_width"); ok {
		outerWidth := int64(v.(int))
		params.Data.OuterWidth = &outerWidth
	}

	if v, ok := d.GetOk("outer_depth"); ok {
		outerDepth := int64(v.(int))
		params.Data.OuterDepth = &outerDepth
	}

	if v, ok := d.GetOk("outer_unit"); ok {
		params.Data.OuterUnit = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Dcim.DcimRacksCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create rack: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceDcimRackRead(ctx, d, m)

	return diags
}

func resourceDcimRackRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	rackID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimRacksReadParams{
		Context: ctx,
		ID:      rackID,
	}

	resp, err := c.Dcim.DcimRacksRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get rack: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("site_id", resp.Payload.Site.ID)
	d.Set("desc_units", resp.Payload.DescUnits)
	d.Set("u_height", resp.Payload.UHeight)

	if resp.Payload.FacilityID != nil {
		d.Set("facility", resp.Payload.FacilityID)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Role != nil {
		d.Set("role_id", resp.Payload.Role.ID)
	}

	if resp.Payload.Serial != "" {
		d.Set("serial", resp.Payload.Serial)
	}

	if resp.Payload.AssetTag != nil {
		d.Set("asset_tag", resp.Payload.AssetTag)
	}

	if resp.Payload.Type != nil {
		d.Set("type", resp.Payload.Type.Value)
	}

	if resp.Payload.Width != nil {
		d.Set("width", resp.Payload.Width.Value)
	}

	if resp.Payload.OuterWidth != nil {
		d.Set("outer_width", resp.Payload.OuterWidth)
	}

	if resp.Payload.OuterDepth != nil {
		d.Set("outer_depth", resp.Payload.OuterDepth)
	}

	if resp.Payload.OuterUnit != nil {
		d.Set("outer_unit", resp.Payload.OuterUnit.Value)
	}

	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceDcimRackUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	rackID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	siteID := int64(d.Get("site_id").(int))

	params := &dcim.DcimRacksPartialUpdateParams{
		Context: ctx,
		ID:      rackID,
	}

	params.Data = &models.WritableRack{
		Name: &name,
		Site: &siteID,
	}

	if d.HasChange("facility") {
		facility := d.Get("facility").(string)
		params.Data.FacilityID = &facility
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		params.Data.Site = &siteID
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("role_id") {
		roleID := int64(d.Get("role_id").(int))
		params.Data.Role = &roleID
	}

	if d.HasChange("serial") {
		params.Data.Serial = d.Get("serial").(string)
	}

	if d.HasChange("asset_tag") {
		aseetTag := d.Get("asset_tag").(string)
		params.Data.AssetTag = &aseetTag
	}

	if d.HasChange("type") {
		params.Data.Type = d.Get("type").(string)
	}

	if d.HasChange("width") {
		width := int64(d.Get("width").(int))
		params.Data.Width = width
	}

	if d.HasChange("u_height") {
		params.Data.UHeight = int64(d.Get("u_height").(int))
	}

	if d.HasChange("desc_units") {
		params.Data.DescUnits = d.Get("desc_units").(bool)
	}

	if d.HasChange("outer_width") {
		outerWidth := int64(d.Get("outer_width").(int))
		params.Data.OuterWidth = &outerWidth
	}

	if d.HasChange("outer_depth") {
		outerDepth := int64(d.Get("outer_depth").(int))
		params.Data.OuterDepth = &outerDepth
	}

	if d.HasChange("outer_unit") {
		params.Data.OuterUnit = d.Get("outer_unit").(string)
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

	_, err = c.Dcim.DcimRacksPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update rack: %v", err)
	}

	return resourceDcimRackRead(ctx, d, m)
}

func resourceDcimRackDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	rackID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimRacksDeleteParams{
		Context: ctx,
		ID:      rackID,
	}

	_, err = c.Dcim.DcimRacksDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete rack: %v", err)
	}

	d.SetId("")

	return diags
}
