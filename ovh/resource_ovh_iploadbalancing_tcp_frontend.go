package ovh

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

type IpLoadbalancingTcpFrontend struct {
	Id            int     `json:"frontendId,omitempty"`
	AllowedSource *string `json:"allowedSource,omitempty"`
	DedicatedIpfo *string `json:"dedicatedIpfo,omitempty"`
	DefaultFarmId *int    `json:"defaultFarmId,omitempty"`
	DefaultSslId  *int    `json:"defaultSslId,omitempty"`
	Disabled      *bool   `json:"disabled,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	Port          int     `json:"port,omitempty"`
	Ssl           *bool   `json:"ssl,omitempty"`
	Zone          string  `json:"zone,omitempty"`
}

func resourceIpLoadbalancingTcpFrontend() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpLoadbalancingTcpFrontendCreate,
		Read:   resourceIpLoadbalancingTcpFrontendRead,
		Update: resourceIpLoadbalancingTcpFrontendUpdate,
		Delete: resourceIpLoadbalancingTcpFrontendDelete,
		Schema: map[string]*schema.Schema{
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowed_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_ipfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_farm_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_ssl_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIpLoadbalancingTcpFrontendCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newFrontend := &IpLoadbalancingTcpFrontend{
		AllowedSource: getNilStringPointer(d.Get("allowed_source").(string)),
		DedicatedIpfo: getNilStringPointer(d.Get("dedicated_ipfo").(string)),
		DefaultFarmId: getNilIntPointer(d.Get("default_farm_id").(int)),
		DefaultSslId:  getNilIntPointer(d.Get("default_ssl_id").(int)),
		Disabled:      getNilBoolPointer(d.Get("disabled").(bool)),
		DisplayName:   getNilStringPointer(d.Get("display_name").(string)),
		Port:          d.Get("port").(int),
		Ssl:           getNilBoolPointer(d.Get("ssl").(bool)),
		Zone:          d.Get("zone").(string),
	}

	service := d.Get("service_name").(string)
	r := &IpLoadbalancingTcpFrontend{}
	endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/frontend", service)

	err := config.OVHClient.Post(endpoint, newFrontend, r)
	if err != nil {
		return fmt.Errorf("calling %s:\n\t %s", endpoint, err.Error())
	}

	//set id
	d.SetId(fmt.Sprintf("%d", r.Id))

	return nil
}

func resourceIpLoadbalancingTcpFrontendRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	service := d.Get("service_name").(string)
	r := &IpLoadbalancingTcpFrontend{}

	endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/frontend/%s", service, d.Id())

	err := config.OVHClient.Get(endpoint, r)
	if err != nil {
		return fmt.Errorf("calling %s :\n\t %q", endpoint, err)
	}
	log.Printf("[DEBUG] Response object from OVH : %v", r)

	if r.AllowedSource != nil {
		d.Set("allowed_source", *r.AllowedSource)
	}
	if r.DedicatedIpfo != nil {
		d.Set("dedicated_ipfo", *r.DedicatedIpfo)
	}
	if r.DefaultFarmId != nil {
		d.Set("default_farm_id", *r.DefaultFarmId)
	}
	if r.DefaultSslId != nil {
		d.Set("default_ssl_id", *r.DefaultSslId)
	}
	if r.Disabled != nil {
		d.Set("disabled", *r.Disabled)
	}
	if r.DisplayName != nil {
		d.Set("display_name", *r.DisplayName)
	}
	if r.Ssl != nil {
		d.Set("ssl", *r.Ssl)
	}
	d.Set("port", r.Port)
	d.Set("zone", r.Zone)

	return nil
}

func resourceIpLoadbalancingTcpFrontendUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	update := &IpLoadbalancingTcpFrontend{
		AllowedSource: getNilStringPointer(d.Get("allowed_source").(string)),
		DedicatedIpfo: getNilStringPointer(d.Get("dedicated_ipfo").(string)),
		DefaultFarmId: getNilIntPointer(d.Get("default_farm_id").(int)),
		DefaultSslId:  getNilIntPointer(d.Get("default_ssl_id").(int)),
		Disabled:      getNilBoolPointer(d.Get("disabled").(bool)),
		DisplayName:   getNilStringPointer(d.Get("display_name").(string)),
		Port:          d.Get("port").(int),
		Ssl:           getNilBoolPointer(d.Get("ssl").(bool)),
		Zone:          d.Get("zone").(string),
	}

	service := d.Get("service_name").(string)
	r := &IpLoadbalancingTcpFrontend{}
	endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/frontend/%s", service, d.Id())
	js, _ := json.Marshal(update)
	log.Printf("[DEBUG] PUT %s : %v", endpoint, string(js))
	err := config.OVHClient.Put(endpoint, update, r)
	if err != nil {
		return fmt.Errorf("calling %s :\n\t %s", endpoint, err.Error())
	}
	return nil
}

func resourceIpLoadbalancingTcpFrontendDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	service := d.Get("service_name").(string)

	r := &IpLoadbalancingTcpFrontend{}
	endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/frontend/%s", service, d.Id())

	err := config.OVHClient.Delete(endpoint, r)
	if err != nil {
		return fmt.Errorf("calling %s :\n\t %s", endpoint, err.Error())
	}

	return nil
}
