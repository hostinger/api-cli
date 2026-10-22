package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Sign in through the browser",
	Long: `Sign in through the browser (OAuth). The tokens are stored locally and
refreshed automatically; commands also sign in on first use when no API token
is configured.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		oauth, err := api.NewOAuthService()
		if err != nil {
			return err
		}
		if _, err := oauth.Login(cmd.Context()); err != nil {
			return fmt.Errorf("sign-in failed: %w", err)
		}
		w := cmd.OutOrStdout()
		fmt.Fprintln(w, "Signed in.")
		if viper.GetString("api_token") != "" {
			fmt.Fprintf(w, "Note: an API token is configured and takes precedence over this sign-in (see '%s auth status').\n", cmd.Root().Name())
		}
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Sign out and remove the stored tokens",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		oauth, err := api.NewOAuthService()
		if err != nil {
			return err
		}
		if err := oauth.Logout(cmd.Context()); err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Signed out.")
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(authLoginCmd, authLogoutCmd)
	oauthStatus = func(w io.Writer) {
		oauth, err := api.NewOAuthService()
		if err != nil {
			fmt.Fprintln(w, "OAuth sign-in: unavailable:", err)
			return
		}
		st := oauth.Status()
		switch {
		case !st.SignedIn:
			fmt.Fprintln(w, "OAuth sign-in: not signed in")
		case st.Expired && st.CanRefresh:
			fmt.Fprintf(w, "OAuth sign-in: signed in (access token expired, refreshes on next use; %s)\n", st.CredPath)
		case st.Expired:
			fmt.Fprintf(w, "OAuth sign-in: expired (%s)\n", st.CredPath)
		case st.ExpiresAt.IsZero():
			fmt.Fprintf(w, "OAuth sign-in: signed in (%s)\n", st.CredPath)
		default:
			fmt.Fprintf(w, "OAuth sign-in: signed in, access token valid until %s (%s)\n", st.ExpiresAt.Local().Format(time.RFC1123), st.CredPath)
		}
	}
}
