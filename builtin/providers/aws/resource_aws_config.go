package aws

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsConfigCreate,
		Read:   resourceAwsConfigRead,
		Update: resourceAwsConfigUpdate,
		Delete: resourceAwsConfigDelete,

		//TO DO: schema for aws config service not implemented yet
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
		},
	}
}

func resourceAwsConfigCreate(d *schema.ResourceData, meta interface{}) error {
	//conn := meta.(*AWSClient).Configconn
	return fmt.Errorf("ConfigCreate Not implemented")

	//return resourceAwsConfigRead(d, meta)
}

func resourceAwsConfigRead(d *schema.ResourceData, meta interface{}) error {
	//conn := meta.(*AWSClient).Configconn

	return fmt.Errorf("resourceAwsConfigRead Not implemented")
	//	return nil
}

func resourceAwsConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	//conn := meta.(*AWSClient).Configconn
	return fmt.Errorf("resourceAwsConfigUpdate Not implemented")

	//	return nil
}

func resourceAwsConfigDelete(d *schema.ResourceData, meta interface{}) error {
	//conn := meta.(*AWSClient).Configconn

	return fmt.Errorf("resourceAwsConfigDelete Not implemented")
	//return nil
}

func configDeclarationSchema() {

}
