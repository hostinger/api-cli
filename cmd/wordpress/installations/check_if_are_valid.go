package installations

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CheckIfAreValidCmd = &cobra.Command{
	Use:   "check-if-are-valid <username>",
	Short: "Check if WordPress installations are valid",
	Long:  "Check whether one or more WordPress installations are valid and working\ncorrectly. Detects broken installations caused by missing files, broken\nplugins, themes and similar issues.\n\nProvide the WordPress installation (software) identifiers in the body. They\ncan be obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield).",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(checkIfAreValidBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCheckIfWordPressInstallationsAreValidV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CheckIfAreValidCmd.Flags().BoolP("force", "", false, "Force fresh validation without cache. Preferable for troubleshooting purposes.")
	CheckIfAreValidCmd.Flags().StringSliceP("software-ids", "", nil, "WordPress installation (software) identifiers to validate.")
	CheckIfAreValidCmd.MarkFlagRequired("software-ids")
}

func checkIfAreValidBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("force") {
		v, _ := cmd.Flags().GetBool("force")
		body["force"] = v
	}
	softwareIdsVal, _ := cmd.Flags().GetStringSlice("software-ids")
	body["software_ids"] = softwareIdsVal
	return body
}
