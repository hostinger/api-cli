package themes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy <username> <domain>",
	Short: "Deploy WordPress theme",
	Long:  "Deploy a WordPress theme from an already uploaded directory.\n\nThis endpoint allows you to deploy a WordPress theme that has been uploaded to the website's directory.\nThe theme can be optionally activated after deployment.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deployBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingDeployWordPressThemeV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeployCmd.Flags().BoolP("is-activated", "", false, "Whether to activate the theme after deployment")
	DeployCmd.Flags().StringP("slug", "", "", "Slug of the theme")
	DeployCmd.Flags().StringP("theme-path", "", "", "Relative path to the theme directory from wp-content/themes")
	DeployCmd.MarkFlagRequired("slug")
	DeployCmd.MarkFlagRequired("theme-path")
}

func deployBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("is-activated") {
		v, _ := cmd.Flags().GetBool("is-activated")
		body["is_activated"] = v
	}
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	themePathVal, _ := cmd.Flags().GetString("theme-path")
	body["theme_path"] = themePathVal
	return body
}
