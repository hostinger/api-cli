package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var SettingsCmd = &cobra.Command{
	Use:   "settings <profile-uuid>",
	Short: "Get profile settings",
	Long:  "Get the profile's saved sender identity and brand details.\n\nRead the default sender before creating a campaign instead of asking the customer, and\nreuse the brand fields when generating email content. Any field can be `null` for a profile\nthat never filled its settings in. This does not check that the sending domain is\nverified, use the connected sending domain endpoint for that.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetProfileSettingsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
