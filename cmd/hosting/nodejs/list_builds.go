package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListBuildsCmd = &cobra.Command{
	Use:   "list-builds <username> <domain>",
	Short: "List NodeJS builds",
	Long:  "Retrieve a paginated list of Node.js build processes for a specific website.\n\nEach build represents a single run of the Node.js build pipeline. Use the `states`\nquery parameter to filter results by build state (pending, running, completed, failed).\nUse the `uuid` from a build to poll its output via the `Get Node.js Build Logs` endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListNodeJSBuildsV1WithResponse(context.TODO(), args[0], args[1], listBuildsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListBuildsCmd.Flags().IntP("page", "", 0, "Page number")
	ListBuildsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListBuildsCmd.Flags().StringSliceP("states", "", nil, "Build states to filter by (one of: pending, running, completed, failed)")
}

func listBuildsParams(cmd *cobra.Command) *client.HostingListNodeJSBuildsV1Params {
	params := &client.HostingListNodeJSBuildsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("states") {
		v, _ := cmd.Flags().GetStringSlice("states")
		es := make([]client.HostingListNodeJSBuildsV1ParamsStates, len(v))
		for i, s := range v {
			es[i] = client.HostingListNodeJSBuildsV1ParamsStates(s)
		}
		params.States = &es
	}
	return params
}
