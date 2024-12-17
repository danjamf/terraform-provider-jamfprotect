package main

import (
	"github.com/danjamf/terraform-provider-jamfprotect/endpoints/preventlists"
	auth "github.com/danjamf/terraform-provider-jamfprotect/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Provider function to define the provider's schema and configuration
func Provider() *schema.Provider {
	return &schema.Provider{
		// Define the provider schema here
		Schema: map[string]*schema.Schema{
			"domainname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Jamf Protect endpoint is your Jamf Protect tenant",
			},
			"clientid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The clientID created for authentication.",
			},
			"clientpassword": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The client password (sic) used for authentication.",
			},
		},

		// Define resources and data sources the provider supports
		ResourcesMap: map[string]*schema.Resource{
			"jamfprotect_preventlist": preventlists.ResourcePreventlists(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"jamfprotect_preventlist": preventlists.DataSourcePreventlists(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	println("hello test test test")
	err := auth.Authenticate(d.Get("domainname").(string), d.Get("clientid").(string), d.Get("clientpassword").(string))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func main() {
	// Start the provider plugin. Do not assign the return value.
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
