package php

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <username> <domain>",
	Short: "Get PHP details",
	Long:  "Returns the full PHP configuration for the website: current version, available versions\n(supported and unsupported), enabled/disabled extensions, options with their current value,\ndefault, type and the plan limit (`max`), and conflicting extension groups.\n\nUse it to check the current PHP setup before updating the version, extensions or options.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetPHPDetailsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
