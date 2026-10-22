package tags

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateOrFindCmd = &cobra.Command{
	Use:               "create-or-find <profile-uuid>",
	Short:             "Create or find tags",
	Long:              "Create tags in a profile.\n\nNames that already exist in the profile are not duplicated: the existing tag is returned\ninstead, so the call is safe to repeat. Every tag in the request is returned, whether it\nwas created now or already existed.",
	Example:           "  hostinger reach tags create-or-find 550e8400-e09b-41d4-a716-400055000000 --names Newsletter,VIP,Promo",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createOrFindBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateOrFindTagsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(CreateOrFindCmd, utils.StringArg("profile-uuid"))
	CreateOrFindCmd.Flags().StringSliceP("names", "", nil, "")
	CreateOrFindCmd.MarkFlagRequired("names")
	CreateOrFindCmd.RegisterFlagCompletionFunc("names", cobra.NoFileCompletions)
}

func createOrFindBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	namesVal, _ := cmd.Flags().GetStringSlice("names")
	body["names"] = namesVal
	return body, nil
}
