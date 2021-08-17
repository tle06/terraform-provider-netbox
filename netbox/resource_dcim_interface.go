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

func resourceDcimInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimInterfaceCreate,
		ReadContext:   resourceDcimInterfaceRead,
		UpdateContext: resourceDcimInterfaceUpdate,
		DeleteContext: resourceDcimInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.InterfaceTypeValueVirtual,
					models.InterfaceTypeValueLag,
					models.InterfaceTypeValueNr100baseTx,
					models.InterfaceTypeValueNr1000baset,
					models.InterfaceTypeValueNr25gbaset,
					models.InterfaceTypeValueNr5gbaset,
					models.InterfaceTypeValueNr10gbaset,
					models.InterfaceTypeValueNr10gbaseCx4,
					models.InterfaceTypeValueNr1000basexGbic,
					models.InterfaceTypeValueNr1000basexSfp,
					models.InterfaceTypeValueNr10gbasexSfpp,
					models.InterfaceTypeValueNr10gbasexXfp,
					models.InterfaceTypeValueNr10gbasexXenpak,
					models.InterfaceTypeValueNr10gbasexX2,
					models.InterfaceTypeValueNr25gbasexSfp28,
					models.InterfaceTypeValueNr40gbasexQsfpp,
					models.InterfaceTypeValueNr50gbasexSfp28,
					models.InterfaceTypeValueNr100gbasexCfp,
					models.InterfaceTypeValueNr100gbasexCfp2,
					models.InterfaceTypeValueNr200gbasexCfp2,
					models.InterfaceTypeValueNr100gbasexCfp4,
					models.InterfaceTypeValueNr100gbasexCpak,
					models.InterfaceTypeValueNr100gbasexQsfp28,
					models.InterfaceTypeValueNr200gbasexQsfp56,
					models.InterfaceTypeValueNr400gbasexQsfpdd,
					models.InterfaceTypeValueNr400gbasexOsfp,
					models.InterfaceTypeValueIeee80211a,
					models.InterfaceTypeValueIeee80211g,
					models.InterfaceTypeValueIeee80211n,
					models.InterfaceTypeValueIeee80211ac,
					models.InterfaceTypeValueIeee80211ad,
					models.InterfaceTypeValueIeee80211ax,
					models.InterfaceTypeValueGsm,
					models.InterfaceTypeValueCdma,
					models.InterfaceTypeValueLte,
					models.InterfaceTypeValueSonetOc3,
					models.InterfaceTypeValueSonetOc12,
					models.InterfaceTypeValueSonetOc48,
					models.InterfaceTypeValueSonetOc192,
					models.InterfaceTypeValueSonetOc768,
					models.InterfaceTypeValueSonetOc1920,
					models.InterfaceTypeValueSonetOc3840,
					models.InterfaceTypeValueNr1gfcSfp,
					models.InterfaceTypeValueNr2gfcSfp,
					models.InterfaceTypeValueNr4gfcSfp,
					models.InterfaceTypeValueNr8gfcSfpp,
					models.InterfaceTypeValueNr16gfcSfpp,
					models.InterfaceTypeValueNr32gfcSfp28,
					models.InterfaceTypeValueNr128gfcSfp28,
					models.InterfaceTypeValueInfinibandSdr,
					models.InterfaceTypeValueInfinibandDdr,
					models.InterfaceTypeValueInfinibandQdr,
					models.InterfaceTypeValueInfinibandFdr10,
					models.InterfaceTypeValueInfinibandFdr,
					models.InterfaceTypeValueInfinibandEdr,
					models.InterfaceTypeValueInfinibandHdr,
					models.InterfaceTypeValueInfinibandNdr,
					models.InterfaceTypeValueInfinibandXdr,
					models.InterfaceTypeValueT1,
					models.InterfaceTypeValueE1,
					models.InterfaceTypeValueT3,
					models.InterfaceTypeValueE3,
					models.InterfaceTypeValueCiscoStackwise,
					models.InterfaceTypeValueCiscoStackwisePlus,
					models.InterfaceTypeValueCiscoFlexstack,
					models.InterfaceTypeValueCiscoFlexstackPlus,
					models.InterfaceTypeValueJuniperVcp,
					models.InterfaceTypeValueExtremeSummitstack,
					models.InterfaceTypeValueExtremeSummitstack128,
					models.InterfaceTypeValueExtremeSummitstack256,
					models.InterfaceTypeValueExtremeSummitstack512,
					models.InterfaceTypeValueOther,
				}),
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"connection_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"management_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.InterfaceModeValueAccess,
					models.InterfaceModeValueTagged,
					models.InterfaceModeValueTaggedAll,
				}),
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tagged_vlan": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"untagged_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"mtu": {
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
		},
	}
}

func resourceDcimInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	interfaceID := int64(d.Get("device_id").(int))
	interfaceType := d.Get("type").(string)
	name := d.Get("name").(string)

	params := &dcim.DcimInterfacesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableInterface{
		Device:      &interfaceID,
		Type:        &interfaceType,
		Name:        &name,
		Tags:        expandTags(d.Get("tags").([]interface{})),
		TaggedVlans: expandTaggedVlans(d.Get("tagged_vlan").([]interface{})),
	}

	if v, ok := d.GetOk("connection_status"); ok {
		connectionStatus := v.(bool)
		params.Data.ConnectionStatus = &connectionStatus
	}

	if v, ok := d.GetOk("enabled"); ok {
		params.Data.Enabled = v.(bool)
	}

	if v, ok := d.GetOk("management_only"); ok {
		params.Data.MgmtOnly = v.(bool)
	}

	if v, ok := d.GetOk("label"); ok {
		params.Data.Label = v.(string)
	}

	if v, ok := d.GetOk("mac_address"); ok {
		macAddress := v.(string)
		params.Data.MacAddress = &macAddress
	}

	if v, ok := d.GetOk("mode"); ok {
		params.Data.Mode = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("untagged_vlan_id"); ok {
		untaggedVlan := int64(v.(int))
		params.Data.UntaggedVlan = &untaggedVlan
	}
	if v, ok := d.GetOk("mtu"); ok {
		mtu := int64(v.(int))
		params.Data.Mtu = &mtu
	}

	resp, err := c.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create interface: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceDcimInterfaceRead(ctx, d, m)

	return diags
}

func resourceDcimInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	interfaceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimInterfacesReadParams{
		Context: ctx,
		ID:      interfaceID,
	}

	resp, err := c.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get interface: %v", err)
	}

	d.Set("device_id", resp.Payload.Device.ID)
	d.Set("type", resp.Payload.Type.Value)
	d.Set("name", resp.Payload.Name)

	if resp.Payload.ConnectionStatus != nil {
		d.Set("connection_status", resp.Payload.ConnectionStatus.Value)
	}

	if resp.Payload.Label != "" {
		d.Set("label", resp.Payload.Label)
	}

	if resp.Payload.MacAddress != nil {
		d.Set("mac_address", resp.Payload.MacAddress)
	}

	if resp.Payload.Mode != nil {
		d.Set("mode", resp.Payload.Mode.Value)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.UntaggedVlan != nil {
		d.Set("untagged_vlan_id", resp.Payload.UntaggedVlan.ID)
	}

	if resp.Payload.Mtu != nil {
		d.Set("mtu", resp.Payload.Mtu)
	}

	d.Set("enabled", resp.Payload.Enabled)
	d.Set("management_only", resp.Payload.MgmtOnly)
	d.Set("tagged_vlan", flattenTaggedVlans(resp.Payload.TaggedVlans))

	d.Set("tags", flattenTags(resp.Payload.Tags))

	return diags
}

func resourceDcimInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	interfaceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	deviceID := int64(d.Get("device_id").(int))
	interfaceType := d.Get("type").(string)
	name := d.Get("name").(string)

	params := &dcim.DcimInterfacesPartialUpdateParams{
		Context: ctx,
		ID:      interfaceID,
	}

	params.Data = &models.WritableInterface{
		Device:      &deviceID,
		Type:        &interfaceType,
		Name:        &name,
		TaggedVlans: expandTaggedVlans(d.Get("tagged_vlan").([]interface{})),
	}

	if d.HasChange("connection_status") {
		connectionStatus := d.Get("connection_status").(bool)
		params.Data.ConnectionStatus = &connectionStatus
	}

	if d.HasChange("enabled") {
		params.Data.Enabled = d.Get("enabled").(bool)
	}

	if d.HasChange("management_only") {
		params.Data.MgmtOnly = d.Get("management_only").(bool)
	}

	if d.HasChange("label") {
		params.Data.Label = d.Get("label").(string)
	}

	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		params.Data.MacAddress = &macAddress
	}

	if d.HasChange("mode") {
		params.Data.Mode = d.Get("mode").(string)
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("untagged_vlan_id") {
		untaggedVlan := int64(d.Get("untagged_vlan_id").(int))
		params.Data.UntaggedVlan = &untaggedVlan
	}

	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		params.Data.Mtu = &mtu
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	_, err = c.Dcim.DcimInterfacesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update interface: %v", err)
	}

	return resourceDcimInterfaceRead(ctx, d, m)
}

func resourceDcimInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	deviceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimInterfacesDeleteParams{
		Context: ctx,
		ID:      deviceID,
	}

	_, err = c.Dcim.DcimInterfacesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete interface: %v", err)
	}

	d.SetId("")

	return diags
}

func expandTaggedVlans(input []interface{}) []int64 {
	if len(input) == 0 {
		return nil
	}

	results := make([]int64, 0)

	for _, item := range input {
		value := item.(int)
		results = append(results, int64(value))
	}

	return results
}

func flattenTaggedVlans(input []*models.NestedVLAN) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["name"] = item.Name
		values["vid"] = item.Vid
		values["DisplayName"] = item.DisplayName

		result = append(result, values["id"])
	}

	return result
}
