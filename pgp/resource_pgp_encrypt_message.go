package pgp

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePGPEncryptMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourcePGPEncryptMessageCreate,
		Read:   resourcePGPEncryptMessageRead,
		Delete: resourcePGPEncryptMessageDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"public_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePGPEncryptMessageCreate(d *schema.ResourceData, m interface{}) error {
	message := d.Get("content").(string)
	public_key := d.Get("public_key").(string)

	// Create public key entity
	publicKeyPacket, _ := getPublicKeyPacket([]byte(public_key))
	pubEntity, _ := createEntityFromKeys(publicKeyPacket, nil)

	// Encrypt message
	encrypted, _ := Encrypt(pubEntity, []byte(message))

	d.Set("result", encrypted)
	d.SetId(sha256sum(encrypted))

	return nil
}

func resourcePGPEncryptMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePGPEncryptMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
