package websites

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <domain>",
	Short: "Delete website",
	Long:  "Permanently deletes a website and all of its data. This action is destructive\nand cannot be undone. Always ask the user for explicit confirmation before\ncalling this endpoint.\n\nAll website files, databases and related configuration will be removed.\nThe hosting plan itself is kept, so a new website can be created on it afterwards.\n\nThe confirm field must be boolean true, otherwise the request is rejected.\n\nSupported websites: main and addon domain websites on web hosting plans, and\nWebsite Builder websites. Parked domains and subdomains cannot be deleted with\nthis endpoint. The domain must be the exact website domain, not a preview\ndomain or an alias.\n\nReturns 404 when the domain does not exist or does not belong to the\nauthenticated client.\n\nWebsite removal is processed asynchronously and can take a few minutes to\ncomplete. The response returns before the removal finishes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deleteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingDeleteWebsiteV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteCmd.Flags().BoolP("confirm", "", false, "Must be boolean true to confirm the permanent deletion of the website. (one of: true)")
	DeleteCmd.MarkFlagRequired("confirm")
}

func deleteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	confirmVal, _ := cmd.Flags().GetBool("confirm")
	body["confirm"] = confirmVal
	return body
}
