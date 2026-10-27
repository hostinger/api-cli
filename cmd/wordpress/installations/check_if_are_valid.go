package installations

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CheckIfAreValidCmd = &cobra.Command{
	Use:     "check-if-are-valid [username]",
	Short:   "Check if WordPress installations are valid",
	Long:    "Check whether one or more WordPress installations are valid and working\ncorrectly. Detects broken installations caused by missing files, broken\nplugins, themes and similar issues.\n\nProvide the WordPress installation (software) identifiers in the body. They\ncan be obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield).\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress installations check-if-are-valid u123456789 --software-ids 123,456",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "username",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, args, err := site.ResolveUsername(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := checkIfAreValidBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCheckIfWordPressInstallationsAreValidV1WithBodyWithResponse(ctx, username, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CheckIfAreValidCmd, 1, utils.StringArg("username"))
	CheckIfAreValidCmd.Flags().BoolP("force", "", false, "Force fresh validation without cache. Preferable for troubleshooting purposes.")
	CheckIfAreValidCmd.Flags().StringSliceP("software-ids", "", nil, "WordPress installation (software) identifiers to validate.")
	CheckIfAreValidCmd.MarkFlagRequired("software-ids")
	CheckIfAreValidCmd.RegisterFlagCompletionFunc("force", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	CheckIfAreValidCmd.RegisterFlagCompletionFunc("software-ids", cobra.NoFileCompletions)
}

func checkIfAreValidBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("force") {
		v, _ := cmd.Flags().GetBool("force")
		body["force"] = v
	}
	softwareIdsVal, _ := cmd.Flags().GetStringSlice("software-ids")
	body["software_ids"] = softwareIdsVal
	return body, nil
}
