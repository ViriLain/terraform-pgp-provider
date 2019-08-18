package pgp

import (
	"crypto/sha256"
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
	message := d.Get("content")
	public_key := d.Get("private_key")

	// Create private key entity
	privEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))

	// Decrypt message
	decrypted, _ := pgp.Decrypt(privEntity, encrypted)

	return nil
}

func resourcePGPDecryptMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePGPDecryptMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
