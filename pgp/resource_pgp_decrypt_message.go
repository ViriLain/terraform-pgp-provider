package gpg

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func resourcePGPDecryptMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGPGEncryptedMessageCreate,
		Read:   resourceGPGEncryptedMessageRead,
		Delete: resourceGPGEncryptedMessageDelete,

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

func resourcePGPEncryptMessage(d *schema.ResourceData, m interface{}) error {
	message := d.Get("content")
	public_key := d.Get("private_key")

	// Create private key entity
	privEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))

	// Decrypt message
	decrypted, _ := pgp.Decrypt(privEntity, encrypted)

	return nil
}

func resourcePGPEncryptedMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGPGEncryptedMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
