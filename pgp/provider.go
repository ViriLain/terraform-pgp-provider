package pgp

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"pgp_encrypt_message": resourcePGPEncryptMessage(),
			"pgp_decrypt_message": resourcePGPDecryptMessage(),
		},
	}
}
