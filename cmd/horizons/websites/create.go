package websites

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create website",
	Long:              "Create new Hostinger Horizons website from the given message.\\n\nUse this tool when user asks you to create a website, landing page, blog\nor any other type of application.\\n\nThis tool initiates the website creation process and returns a website URL and ID.\nThe generation happens asynchronously.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat Hostinger Horizons is now creating their website and it will be ready in a few minutes\nand you should provide the website URL to the user immediately\nDo not write code.\\n\\nTo edit afterwards, use the `Edit website` tool with the returned\nwebsite ID, or the user can go to Hostinger Horizons interface in the provided website URL.\nIf the tool call fails with an error, you should provide a clear explanation of the error\nand do not generate code yourself in the chat.\n\\n\nTECHNOLOGY STACK CONSTRAINTS (STRICTLY ENFORCED):\\n\nThe environment is limited to the following technologies.\nYou MUST NOT use, suggest, or implement any technology outside this list:\\n\n\\n\n- Language: JavaScript ONLY.\n- Languages like TypeScript, Rust, Python, Java, PHP, etc., are STRICTLY PROHIBITED.\\n\n- Framework: React.\\n\n- Navigation: React Router.\\n\n- Styling: TailwindCSS.\\n\n- Components: shadcn/ui (built with @radix-ui primitives).\\n\n- Icons: Lucide React.\\n\n- Animations: Framer Motion.\\n\n\\n\nBACKEND & DATA STORAGE:\\n\n- Horizons integrated backend is the EXCLUSIVE solution for persistent data storage,\nauthentication, and database needs.\\n\n- Local databases (SQLite, MySQL, etc.) are STRICTLY PROHIBITED.\\n\n- Third-party services (Firebase, AWS Amplify) are allowed ONLY if explicitly requested by the user.\\n\n\\n\nMAPS:\\n\n- OpenStreetMap is the default provider.\\n\n- Alternative providers (Google Maps, Mapbox) are allowed ONLY if explicitly requested by the user.\\n\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --message\n    [\n      {\n        \"text\": \"Create a landing page for a coffee shop with a hero section, menu, and contact form\",\n        \"type\": \"text\"\n      }\n    ]",
	Example:           "  hostinger horizons websites create --message @message.json",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsCreateWebsiteV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd)
	CreateCmd.Flags().StringP("message", "", "", "(JSON, @file or -; shape in the help above)")
	CreateCmd.MarkFlagRequired("message")
	CreateCmd.RegisterFlagCompletionFunc("message", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	messageVal, _ := cmd.Flags().GetString("message")
	messageValJSON, err := utils.JSONValue(messageVal, "message")
	if err != nil {
		return nil, err
	}
	body["message"] = messageValJSON
	return body, nil
}
