package netbox

import (
	"context"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/circuits"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCircuitsCircuit() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCircuitsCircuitCreate,
		ReadContext:   resourceCircuitsCircuitRead,
		UpdateContext: resourceCircuitsCircuitUpdate,
		DeleteContext: resourceCircuitsCircuitDelete,

		Schema: map[string]*schema.Schema{
			"cid": {
				Type:     schema.TypeString,
				Required: true,
			},

			"type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"provider_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.CircuitStatusValueActive,
					models.CircuitStatusValueDecommissioned,
					models.CircuitStatusValueDeprovisioning,
					models.CircuitStatusValueOffline,
					models.CircuitStatusValuePlanned,
					models.CircuitStatusValueProvisioning,
				}),
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"commit_rate": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},

			"comments": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},
			"install_date": {
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

func resourceCircuitsCircuitCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	cid := d.Get("cid").(string)
	circuitType := int64(d.Get("type_id").(int))
	providerID := int64(d.Get("provider_id").(int))

	params := &circuits.CircuitsCircuitsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableCircuit{
		Cid:      &cid,
		Type:     &circuitType,
		Provider: &providerID,
		Tags:     expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("commit_rate"); ok {
		commitRate := int64(v.(int))
		params.Data.CommitRate = &commitRate
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("install_date"); ok {
		installDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			installDateConverted := strfmt.Date(installDate)
			params.Data.InstallDate = &installDateConverted
		}
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Circuits.CircuitsCircuitsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create circuit: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceCircuitsCircuitRead(ctx, d, m)

	return diags
}

func resourceCircuitsCircuitRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &circuits.CircuitsCircuitsReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Circuits.CircuitsCircuitsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get circuit: %v", err)
	}

	d.Set("cid", resp.Payload.Cid)
	d.Set("type_id", resp.Payload.Type.ID)
	d.Set("provider_id", resp.Payload.Provider.ID)

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}
	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}
	if resp.Payload.CommitRate != nil {
		d.Set("commit_rate", resp.Payload.CommitRate)
	}

	if resp.Payload.InstallDate != nil {
		d.Set("install_date", resp.Payload.InstallDate)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceCircuitsCircuitUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	cid := d.Get("cid").(string)
	circuitType := int64(d.Get("type_id").(int))
	providerID := int64(d.Get("provider_id").(int))

	params := &circuits.CircuitsCircuitsPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.WritableCircuit{
		Cid:      &cid,
		Type:     &circuitType,
		Provider: &providerID,
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("commit_rate") {
		commitRate := int64(d.Get("commit_rate").(int))
		params.Data.CommitRate = &commitRate
	}

	if d.HasChange("install_date") {
		installDate, err := time.Parse(time.RFC3339, d.Get("install_date").(string))
		if err != nil {
			installDateConverted := strfmt.Date(installDate)
			params.Data.InstallDate = &installDateConverted
		}
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

	_, err = c.Circuits.CircuitsCircuitsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update circuit: %v", err)
	}

	return resourceCircuitsCircuitRead(ctx, d, m)
}

func resourceCircuitsCircuitDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &circuits.CircuitsCircuitsDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Circuits.CircuitsCircuitsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete cicruit: %v", err)
	}

	d.SetId("")

	return diags
}
