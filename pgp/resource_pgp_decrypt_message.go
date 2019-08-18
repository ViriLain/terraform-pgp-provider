package pgp

import (
	"crypto/sha256"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePGPDecryptMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourcePGPDecryptMessageCreate,
		Read:   resourcePGPDecryptMessageRead,
		Delete: resourcePGPDecryptMessageDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"public_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePGPDecryptMessageCreate(d *schema.ResourceData, m interface{}) error {
	message := d.Get("content").(string)
	private_key := d.Get("private_key").(string)
	public_key := d.Get("public_key").(string)

	// Create private key entity
	privEntity, _ := GetEntity([]byte(public_key), []byte(private_key))

	// Decrypt message
	decrypted, _ := Decrypt(privEntity, []byte(message))

	d.Set("result", decrypted)

	return nil
}

func resourcePGPDecryptMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePGPDecryptMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func sha256sum(data interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data.(string))))
}
