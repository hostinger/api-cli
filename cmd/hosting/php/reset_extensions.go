package php

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ResetExtensionsCmd = &cobra.Command{
	Use:   "reset-extensions <username> <domain>",
	Short: "Reset PHP extensions",
	Long:  "Resets all PHP extensions of the website to their default state.\n\nUse it to recover from extension conflicts or restore the original configuration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingResetPHPExtensionsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
