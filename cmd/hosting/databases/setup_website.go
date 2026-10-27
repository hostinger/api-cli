package databases

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

var SetupWebsiteCmd = &cobra.Command{
	Use:     "setup-website [domain]",
	Short:   "Setup website database",
	Long:    "Creates a new MySQL database for the website and writes its connection details into the\nwebsite's environment variables, then restarts the application. The platform generates the\npassword (and the database name and user, unless supplied). The password is never returned;\nthe application reads it from the environment.\n\nWritten variables: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` and\n`DATABASE_URL` (`mysql://user:password@host:port/name`, user and password percent-encoded).\nExisting variables are kept. If the website already has any variable with one of these\nnames the call fails with 422 and nothing is created; the `Replace Node.js environment\nvariables` endpoint removes them.\n\nAfter this call the variables are ordinary environment variables: the\n`Replace Node.js environment variables` endpoint changes or removes them like any other.\n\nA restart is enough for apps that read environment variables at process start, such as\nExpress or NestJS. Frameworks that bake variables into the build output (Next.js,\n`NEXT_PUBLIC_*`) see the new values only after a fresh build (`Start Node.js build` endpoint).\n\nA password in the request is ignored; the platform always generates it. The optional `name`\nand `user` are identifiers, not secrets.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting databases setup-website mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
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
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := setupWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingSetupWebsiteDatabaseV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(SetupWebsiteCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	SetupWebsiteCmd.Flags().StringP("name", "", "", "Optional database name. Generated when omitted. Letters, digits and underscores; must not start with an underscore. Up to 14 characters without the account username prefix (`u123456789_`), which is added automatically when missing. With the prefix the full name is 12 to 25 characters.")
	SetupWebsiteCmd.Flags().StringP("user", "", "", "Optional database user. Generated when omitted. Letters, digits and underscores; must not start with an underscore. Up to 14 characters without the account username prefix (`u123456789_`), which is added automatically when missing. With the prefix the full user is 12 to 25 characters.")
	SetupWebsiteCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	SetupWebsiteCmd.RegisterFlagCompletionFunc("user", cobra.NoFileCompletions)
}

func setupWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("user") {
		v, _ := cmd.Flags().GetString("user")
		body["user"] = v
	}
	return body, nil
}
