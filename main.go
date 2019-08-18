package pgp

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/littleboytimmy/terraform-pgp-provider/pgp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pgp.Provider})
}
