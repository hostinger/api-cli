package aliases

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateAliasCmd = &cobra.Command{
	Use:               "create-alias <mailbox-id>",
	Short:             "Create alias",
	Long:              "Create an alias for the given mailbox. The alias address is formed\nfrom the given local part and the domain of the mailbox. Messages\nsent to the alias are delivered to the mailbox.",
	Example:           "  hostinger mail aliases create-alias AC1a2b3c4d5e6f7g --local-part info",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createAliasBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateAliasV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateAliasCmd, utils.StringArg("mailbox-id"))
	CreateAliasCmd.Flags().StringP("local-part", "", "", "Local part of the alias address (the part before the @). The domain is taken from the mailbox. Case-insensitive and stored lowercase; must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.")
	CreateAliasCmd.MarkFlagRequired("local-part")
	CreateAliasCmd.RegisterFlagCompletionFunc("local-part", cobra.NoFileCompletions)
}

func createAliasBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	localPartVal, _ := cmd.Flags().GetString("local-part")
	body["local_part"] = localPartVal
	return body, nil
}
