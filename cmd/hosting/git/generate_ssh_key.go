package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GenerateSshKeyCmd = &cobra.Command{
	Use:   "generate-ssh-key <username>",
	Short: "Generate Git SSH key",
	Long:  "Creates the SSH key pair of the hosting account and returns the public key. When the account already\nhas a key, returns that key unchanged. One key serves every website of the account; add the public\nkey to a private repository as a deploy key before deploying it.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGenerateGitSSHKeyV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
