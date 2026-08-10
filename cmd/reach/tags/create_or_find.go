package tags

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateOrFindCmd = &cobra.Command{
	Use:   "create-or-find <profile-uuid>",
	Short: "Create or find tags",
	Long:  "Create tags in a profile.\n\nNames that already exist in the profile are not duplicated: the existing tag is returned\ninstead, so the call is safe to repeat. Every tag in the request is returned, whether it\nwas created now or already existed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createOrFindBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateOrFindTagsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateOrFindCmd.Flags().StringSliceP("names", "", nil, "")
	CreateOrFindCmd.MarkFlagRequired("names")
}

func createOrFindBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	namesVal, _ := cmd.Flags().GetStringSlice("names")
	body["names"] = namesVal
	return body
}
