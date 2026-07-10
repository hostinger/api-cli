package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RestartApplicationCmd = &cobra.Command{
	Use:   "restart-application <username> <domain>",
	Short: "Restart Node.js application",
	Long:  "Restarts the Node.js server process for the website. Does not rebuild or redeploy the\napplication. Use it to apply environment or configuration changes, or to recover a hung\napplication.\n\nOnly applicable to server-side applications (Express, Next.js, NestJS, etc.). Static\nfront-end apps (React, Vue, Vite) have no persistent server process, so restarting them\nhas no effect. Returns success even when the website has no server process to restart.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingRestartNodeJsApplicationV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
