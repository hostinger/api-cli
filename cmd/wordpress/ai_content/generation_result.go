package ai_content

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GenerationResultCmd = &cobra.Command{
	Use:   "generation-result <username> <software> <generation_id>",
	Short: "Get generation result",
	Long:  "Get the result of an AI content generation started for the specified WordPress\ninstallation, such as an image generation started with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/content/images/generate.\n\nWhile the generation is still running only `status: in-progress` is returned; poll every\nfew seconds. Once `completed`, `result` holds the generated content shaped according to\n`type`: temporary image URLs together with the time they expire for `image-generation`, or\nthe generated website document for `website-generation` (started from the Hostinger\nWordPress plugin). A `failed` generation has no `result`. Unknown or expired generation\nIDs return 404.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"website-generation", "image-generation"})
		r, err := api.Request().HostingGetGenerationResultV1WithResponse(context.TODO(), args[0], args[1], uuid.MustParse(args[2]), generationResultParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	GenerationResultCmd.Flags().StringP("type", "", "", "Generation type to look up. When omitted, every generation type is checked. (one of: website-generation, image-generation)")
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
