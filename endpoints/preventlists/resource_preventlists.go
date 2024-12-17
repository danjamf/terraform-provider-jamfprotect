package preventlists

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource definition (this is where you'll manage state, create/update resources)
func ResourcePreventlists() *schema.Resource {
	return &schema.Resource{
		Create: resourceExampleCreate,
		Read:   resourceExampleRead,
		Update: resourceExampleUpdate,
		Delete: resourceExampleDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceExampleCreate(d *schema.ResourceData, m interface{}) error {
	// Create the resource (for example, make an API call to create the resource)
	name := d.Get("name").(string)
	fmt.Printf("Creating resource with name: %s\n", name)

	// Set the resource ID
	d.SetId(name) // The ID is used by Terraform to track the resource

	return resourceExampleRead(d, m)
}

func resourceExampleRead(d *schema.ResourceData, m interface{}) error {
	// Read the resource data (e.g., make an API call to fetch resource details)
	name := d.Id()
	fmt.Printf("Reading resource with ID: %s\n", name)

	// Return current state (for now, just set the name as the resource ID)
	d.Set("name", name)
	return nil
}

func resourceExampleUpdate(d *schema.ResourceData, m interface{}) error {
	// Update the resource (e.g., make an API call to update the resource)
	name := d.Get("name").(string)
	fmt.Printf("Updating resource with name: %s\n", name)

	// Update logic here (if needed)

	return resourceExampleRead(d, m)
}

func resourceExampleDelete(d *schema.ResourceData, m interface{}) error {
	// Delete the resource (e.g., make an API call to delete the resource)
	name := d.Id()
	fmt.Printf("Deleting resource with ID: %s\n", name)

	// Delete logic here (if needed)

	// Remove resource from state
	d.SetId("")
	return nil
}
