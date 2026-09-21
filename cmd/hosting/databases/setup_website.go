package databases

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var SetupWebsiteCmd = &cobra.Command{
	Use:   "setup-website <username> <domain>",
	Short: "Setup website database",
	Long:  "Creates a new MySQL database for the website and writes its connection details into the\nwebsite's environment variables, then restarts the application. The platform generates the\npassword (and the database name and user, unless supplied). The password is never returned;\nthe application reads it from the environment.\n\nWritten variables: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` and\n`DATABASE_URL` (`mysql://user:password@host:port/name`, user and password percent-encoded).\nExisting variables are kept. If the website already has any variable with one of these\nnames the call fails with 422 and nothing is created; the `Replace Node.js environment\nvariables` endpoint removes them.\n\nAfter this call the variables are ordinary environment variables: the\n`Replace Node.js environment variables` endpoint changes or removes them like any other.\n\nA restart is enough for apps that read environment variables at process start, such as\nExpress or NestJS. Frameworks that bake variables into the build output (Next.js,\n`NEXT_PUBLIC_*`) see the new values only after a fresh build (`Start Node.js build` endpoint).\n\nA password in the request is ignored; the platform always generates it. The optional `name`\nand `user` are identifiers, not secrets.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setupWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingSetupWebsiteDatabaseV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetupWebsiteCmd.Flags().StringP("name", "", "", "Optional database name. Generated when omitted. Letters, digits and underscores;\nmust not start with an underscore. Up to 14 characters without the account username\nprefix (`u123456789_`), which is added automatically when missing. With the prefix\nthe full name is 12 to 25 characters.")
	SetupWebsiteCmd.Flags().StringP("user", "", "", "Optional database user. Generated when omitted. Letters, digits and underscores;\nmust not start with an underscore. Up to 14 characters without the account username\nprefix (`u123456789_`), which is added automatically when missing. With the prefix\nthe full user is 12 to 25 characters.")
}

func setupWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("user") {
		v, _ := cmd.Flags().GetString("user")
		body["user"] = v
	}
	return body
}
