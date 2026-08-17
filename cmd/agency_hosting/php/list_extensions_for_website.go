package php

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListExtensionsForWebsiteCmd = &cobra.Command{
	Use:   "list-extensions-for-website <website_uid>",
	Short: "List PHP extensions for a website",
	Long:  "Lists every PHP extension available to an Agency Plan website and whether it is currently enabled.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListPHPExtensionsForAWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
