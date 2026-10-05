package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var SshPublicKeyCmd = &cobra.Command{
	Use:   "ssh-public-key <username>",
	Short: "Get Git SSH public key",
	Long:  "Returns the public SSH key of the hosting account. `Deploy website Git repository` uses this key to\nclone and pull over SSH, so a private repository works once the key is added to it as a deploy key\non the Git host. `public_key` is null when the account has no key yet.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetGitSSHPublicKeyV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
