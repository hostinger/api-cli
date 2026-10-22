package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// oauthStatus prints the OAuth sign-in state. It is set by auth_oauth.go,
// which builds without OAuth sign-in do not ship.
var oauthStatus func(w io.Writer)

var AuthCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Sign in and check which credentials are used",
	GroupID: groupSetup,
	Args:    cobra.ArbitraryArgs,
	RunE:    utils.GroupRunE,
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show which credentials the CLI uses",
	Long: `Show which credentials the CLI uses. A static API token (from the
environment or the config file) takes precedence over any sign-in.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		w := cmd.OutOrStdout()
		token := viper.GetString("api_token")
		switch {
		case token == "":
			fmt.Fprintln(w, "API token: not set")
		case os.Getenv(utils.EnvPrefix+"API_TOKEN") != "":
			fmt.Fprintf(w, "API token: %s (from %sAPI_TOKEN, used for all requests)\n", maskSecret(token), utils.EnvPrefix)
		case viper.InConfig("api_token"):
			fmt.Fprintf(w, "API token: %s (from %s, used for all requests)\n", maskSecret(token), viper.ConfigFileUsed())
		default:
			fmt.Fprintf(w, "API token: %s (used for all requests)\n", maskSecret(token))
		}
		if oauthStatus != nil {
			oauthStatus(w)
		} else if token == "" {
			fmt.Fprintf(w, "Set %sAPI_TOKEN or run '%s config set api_token'.\n", utils.EnvPrefix, cmd.Root().Name())
		}
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(authStatusCmd)
}

// maskSecret keeps only the last four characters of a secret.
func maskSecret(s string) string {
	if len(s) <= 8 {
		return "****"
	}
	return "****" + s[len(s)-4:]
}
