package ovh

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// var TestAccIpLoadbalancingTcpFrontendPlan = [][]map[string]interface{}{
// 	{
// 		{"Status": "active", "Address": os.Getenv("OVH_IPLB_SERVER"), "Port": 80, "Weight": 3},
// 		// {"Port": 8080, "Probe": true, "Backup": true},
// 		// {"Port": 8080, "Probe": false, "Backup": false, "Weight": 2},
// 	},
// 	// 	{
// 	// 		{"Status": "inactive", "Address": os.Getenv("OVH_IPLB_SERVER"), "Port": 80},
// 	// 		{"Port": 8080, "ProxyProtocolVersion": "v2", "Ssl": true},
// 	// 		{"Port": 8080, "ProxyProtocolVersion": "v1", "Ssl": true, "Backup": false},
// 	// 		{"Port": 8080, "ProxyProtocolVersion": nil, "Ssl": true, "Backup": true, "Status": "active"},
// 	// 	},
// }

// type TestAccIpLoadbalancingTcpFrontend struct {
// 	Id            int     `json:"frontendId"`
// 	AllowedSource *string `json:"allowedSource"`
// 	DedicatedIpfo *string `json:"dedicatedIpfo"`
// 	DefaultFarmId *int    `json:"defaultFarmId"`
// 	DefaultSslId  *int    `json:"defaultSslId"`
// 	Disabled      *bool   `json:"disabled"`
// 	DisplayName   *string `json:"displayName"`
// 	Port          int     `json:"port"`
// 	Ssl           *bool   `json:"ssl"`
// 	Zone          string  `json:"zone"`
// }

// type TestAccIpLoadbalancingTcpFrontendWrapper struct {
// 	Response *TestAccIpLoadbalancingTcpFrontend
// 	Expected *TestAccIpLoadbalancingTcpFrontend
// }

// func (w *TestAccIpLoadbalancingTcpFrontendWrapper) Config() string {
// 	var config bytes.Buffer
// 	config.WriteString(fmt.Sprintf(`
// 	resource "ovh_iploadbalancing_tcp_frontend" "testacc" {
// 	  service_name = "%s"
// 	  zone = "%s"
// 	  port = "%s"
// 	`, os.Getenv("OVH_IPLB_SERVICE"),
// 		w.Expected.Zone,
// 		w.Expected.Port))

// 	conditionalAttributeString(&config, "allowed_source", *w.Expected.AllowedSource))
// 	conditionalAttributeString(&config, "dedicated_ipfo", *w.Expected.DedicatedIpfo))
// 	conditionalAttributeInt(&config, "default_farm_id", w.Expected.DefaultFarmId)
// 	conditionalAttributeInt(&config, "default_ssl_id", w.Expected.DefaultSslId)
// 	conditionalAttributeBool(&config, "disabled", w.Expected.Disabled)
// 	conditionalAttributeString(&config, "display_name", *w.Expected.DisplayName))
// 	conditionalAttributeBool(&config, "ssl", w.Expected.Ssl)

// 	config.WriteString(`}`)
// 	log.Printf("[DEBUG] config : %s", config.String())
// 	return config.String()
// }

// func (w *TestAccIpLoadbalancingTcpFrontendWrapper) Equals() bool {
// 	AllowedSource *string `json:"allowedSource"`
// 	DedicatedIpfo *string `json:"dedicatedIpfo"`
// 	DefaultFarmId *int    `json:"defaultFarmId"`
// 	DefaultSslId  *int    `json:"defaultSslId"`
// 	Disabled      *bool   `json:"disabled"`
// 	DisplayName   *string `json:"displayName"`
// 	Port          int     `json:"port"`
// 	Ssl           *bool   `json:"ssl"`
// 	Zone          string  `json:"zone"`

// 	if *w.Response.Address == *w.Expected.Address ||
// 		*w.Response.Port == *w.Expected.Port ||
// 		w.Response.ProxyProtocolVersion == w.Expected.ProxyProtocolVersion || *w.Response.ProxyProtocolVersion == *w.Expected.ProxyProtocolVersion ||
// 		w.Response.Chain == w.Expected.Chain || *w.Response.Chain == *w.Expected.Chain ||
// 		w.Response.Weight == w.Expected.Weight || *w.Response.Weight == *w.Expected.Weight ||
// 		w.Response.Probe == w.Expected.Probe ||
// 		*w.Response.Probe == *w.Expected.Probe ||
// 		w.Response.Ssl == w.Expected.Ssl || *w.Response.Ssl == *w.Expected.Ssl ||
// 		w.Response.Backup == w.Expected.Backup || *w.Response.Backup == *w.Expected.Backup ||
// 		*w.Response.Status == *w.Expected.Status {
// 		return true
// 	}
// 	return false
// }

// func (w *TestAccIpLoadbalancingTcpFrontendWrapper) TestStep(c map[string]interface{}) resource.TestStep {
// 	if val, ok := c["Address"]; ok {
// 		w.Expected.Address = getNilStringPointer(val)
// 	}
// 	if val, ok := c["Port"]; ok {
// 		w.Expected.Port = getNilIntPointer(val)
// 	}
// 	if val, ok := c["ProxyProtocolVersion"]; ok {
// 		w.Expected.ProxyProtocolVersion = getNilStringPointer(val)
// 	}
// 	if val, ok := c["Chain"]; ok {
// 		w.Expected.Chain = getNilStringPointer(val)
// 	}
// 	if val, ok := c["Weight"]; ok {
// 		w.Expected.Weight = getNilIntPointer(val)
// 	}
// 	if val, ok := c["Probe"]; ok {
// 		w.Expected.Probe = getNilBoolPointer(val)
// 	}
// 	if val, ok := c["Ssl"]; ok {
// 		w.Expected.Ssl = getNilBoolPointer(val)
// 	}
// 	if val, ok := c["Backup"]; ok {
// 		w.Expected.Backup = getNilBoolPointer(val)
// 	}
// 	if val, ok := c["Status"]; ok {
// 		w.Expected.Status = getNilStringPointer(val)
// 	}

// 	return resource.TestStep{
// 		Config: w.Config(),
// 		Check: resource.ComposeTestCheckFunc(
// 			w.TestCheck(),
// 		),
// 	}
// }

// func (w *TestAccIpLoadbalancingTcpFrontendWrapper) TestCheck() resource.TestCheckFunc {
// 	return func(state *terraform.State) error {
// 		name := "ovh_iploadbalancing_tcp_farm_server.testacc"
// 		resource, ok := state.RootModule().Resources[name]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", name)
// 		}
// 		config := testAccProvider.Meta().(*Config)
// 		endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/farm/%s/server/%s", os.Getenv("OVH_IPLB_SERVICE"), resource.Primary.Attributes["farm_id"], resource.Primary.ID)
// 		err := config.OVHClient.Get(endpoint, w.Response)
// 		if err != nil {
// 			return fmt.Errorf("calling GET %s :\n\t %s", endpoint, err.Error())
// 		}
// 		if !w.Equals() {
// 			return fmt.Errorf("%s %s state differs from expected", name, resource.Primary.ID)
// 		}
// 		return nil
// 	}
// }

// func (w *TestAccIpLoadbalancingTcpFrontendWrapper) TestDestroy(state *terraform.State) error {
// 	leftovers := false
// 	for _, resource := range state.RootModule().Resources {
// 		if resource.Type != "ovh_iploadbalancing_tcp_farm_server" {
// 			continue
// 		}

// 		config := testAccProvider.Meta().(*Config)
// 		endpoint := fmt.Sprintf("/ipLoadbalancing/%s/tcp/farm/%d/server/%s", os.Getenv("OVH_IPLB_SERVICE"), w.Response.FarmId, resource.Primary.ID)
// 		err := config.OVHClient.Get(endpoint, nil)
// 		if err == nil {
// 			leftovers = true
// 		}
// 	}
// 	if leftovers {
// 		return fmt.Errorf("IpLoadbalancing farm still exists")
// 	}
// 	return nil
// }

// func newTestAccIpLoadbalancingTcpFrontendWrapper() *TestAccIpLoadbalancingTcpFrontendWrapper {
// 	return &TestAccIpLoadbalancingTcpFrontendWrapper{
// 		Expected: &TestAccIpLoadbalancingTcpFrontend{ServiceName: os.Getenv("OVH_IPLB_SERVICE")},
// 		Response: &TestAccIpLoadbalancingTcpFrontend{},
// 	}
// }

// func TestAccIpLoadbalancingTcpFrontendBasicCreate(t *testing.T) {
// 	for _, plan := range TestAccIpLoadbalancingTcpFrontendPlan {
// 		w := newTestAccIpLoadbalancingTcpFrontendWrapper()
// 		var steps []resource.TestStep
// 		for _, tcase := range plan {
// 			steps = append(steps, w.TestStep(tcase))
// 		}
// 		resource.Test(t, resource.TestCase{
// 			Providers:    testAccProviders,
// 			CheckDestroy: w.TestDestroy,
// 			Steps:        steps,
// 		})
// 	}
// }
