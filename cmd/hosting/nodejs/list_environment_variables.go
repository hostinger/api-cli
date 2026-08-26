package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListEnvironmentVariablesCmd = &cobra.Command{
	Use:   "list-environment-variables <username> <domain>",
	Short: "List Node.js environment variables",
	Long:  "Lists the Node.js environment variables currently set for the website. Values are always\nmasked as `********` and cannot be read back through this API. Use this endpoint to see\nwhich keys are configured or to verify a change, not to read values.\n\nTo change variables, use the `Replace Node.js environment variables` endpoint. It replaces\nthe whole set, so never copy the masked values from this response into that request; send\nthe full desired set with real values taken from the project `.env` file or the user\nprompt instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListNodeJsEnvironmentVariablesV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
