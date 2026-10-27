package ai_content

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GenerationResultCmd = &cobra.Command{
	Use:     "generation-result [username] <software> <generation_id>",
	Short:   "Get generation result",
	Long:    "Get the result of an AI content generation started for the specified WordPress\ninstallation, such as an image generation started with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/content/images/generate.\n\nWhile the generation is still running only `status: in-progress` is returned; poll every\nfew seconds. Once `completed`, `result` holds the generated content shaped according to\n`type`: temporary image URLs together with the time they expire for `image-generation`, or\nthe generated website document for `website-generation` (started from the Hostinger\nWordPress plugin). A `failed` generation has no `result`. Unknown or expired generation\nIDs return 404.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress ai-content generation-result u123456789 1232456789 9d4f8a1e-2b3c-4d5e-8f6a-7b8c9d0e1f2a",
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
		username, args, err := site.ResolveUsername(cmd, args, 2)
		if err != nil {
			return api.CommandError(err)
		}
		if err := utils.EnumCheck(cmd, "type", []string{"website-generation", "image-generation"}); err != nil {
			return err
		}
		params := generationResultParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetGenerationResultV1WithResponse(ctx, username, args[0], uuid.MustParse(args[1]), params)
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
	utils.RegisterScopedArgs(GenerationResultCmd, 1, utils.StringArg("username"), utils.StringArg("software"), utils.UUIDArg("generation_id"))
	GenerationResultCmd.Flags().StringP("type", "", "", "Generation type to look up. When omitted, every generation type is checked. (one of: website-generation, image-generation)")
	GenerationResultCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"website-generation", "image-generation"}, cobra.ShellCompDirectiveNoFileComp))
}

func generationResultParams(cmd *cobra.Command) *client.HostingGetGenerationResultV1Params {
	params := &client.HostingGetGenerationResultV1Params{}
	if cmd.Flags().Changed("type") {
		v, _ := cmd.Flags().GetString("type")
		e := client.HostingGetGenerationResultV1ParamsType(v)
		params.Type = &e
	}
	return params
}
