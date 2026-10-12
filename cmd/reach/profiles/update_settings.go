package profiles

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateSettingsCmd = &cobra.Command{
	Use:   "update-settings <profile-uuid>",
	Short: "Update profile settings",
	Long:  "Change the profile's default sender and brand details.\n\nSend only the fields to change: every omitted field keeps its value, and `null` or an empty\nstring clears a text field. The default sender is what campaigns created afterwards start\nfrom; existing campaigns keep their own sender. `sender_email` is only the part before the\n`@`, because Reach combines it with the profile's connected sending domain, so an address\noutside that domain cannot be set. Connecting a sending domain is a separate endpoint and\na connected domain cannot be changed here. Returns the settings after the change.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateSettingsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateProfileSettingsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateSettingsCmd.Flags().StringP("about", "", "", "")
	UpdateSettingsCmd.Flags().StringP("accent-color", "", "", "")
	UpdateSettingsCmd.Flags().StringP("background-color", "", "", "")
	UpdateSettingsCmd.Flags().StringP("brand-name", "", "", "")
	UpdateSettingsCmd.Flags().StringP("business-address", "", "", "")
	UpdateSettingsCmd.Flags().StringP("button-color", "", "", "")
	UpdateSettingsCmd.Flags().IntP("button-corners", "", 0, "")
	UpdateSettingsCmd.Flags().StringP("button-text", "", "", "Button text colour.")
	UpdateSettingsCmd.Flags().StringP("font-family", "", "", "")
	UpdateSettingsCmd.Flags().StringP("heading-font-family", "", "", "")
	UpdateSettingsCmd.Flags().StringSliceP("keywords", "", nil, "")
	UpdateSettingsCmd.Flags().StringP("links-color", "", "", "")
	UpdateSettingsCmd.Flags().StringP("logo-url", "", "", "Public image URL. Logo uploads are not supported.")
	UpdateSettingsCmd.Flags().StringP("privacy-policy-url", "", "", "")
	UpdateSettingsCmd.Flags().StringSliceP("selling-points", "", nil, "")
	UpdateSettingsCmd.Flags().StringP("sender-email", "", "", "Part of the default sender address before the `@`, for example `hello`. Never send a\nfull address: Reach adds the profile's connected sending domain itself.")
	UpdateSettingsCmd.Flags().StringP("sender-name", "", "", "Default sender name. Letters, numbers and spaces only.")
	UpdateSettingsCmd.Flags().StringSliceP("social-links", "", nil, "")
	UpdateSettingsCmd.Flags().StringP("socials", "", "", "Social profile links keyed by platform. (JSON)")
	UpdateSettingsCmd.Flags().StringP("text-color", "", "", "")
	UpdateSettingsCmd.Flags().StringSliceP("topics", "", nil, "")
	UpdateSettingsCmd.Flags().StringP("website-url", "", "", "")
}

func updateSettingsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("about") {
		v, _ := cmd.Flags().GetString("about")
		body["about"] = v
	}
	if cmd.Flags().Changed("accent-color") {
		v, _ := cmd.Flags().GetString("accent-color")
		body["accent_color"] = v
	}
	if cmd.Flags().Changed("background-color") {
		v, _ := cmd.Flags().GetString("background-color")
		body["background_color"] = v
	}
	if cmd.Flags().Changed("brand-name") {
		v, _ := cmd.Flags().GetString("brand-name")
		body["brand_name"] = v
	}
	if cmd.Flags().Changed("business-address") {
		v, _ := cmd.Flags().GetString("business-address")
		body["business_address"] = v
	}
	if cmd.Flags().Changed("button-color") {
		v, _ := cmd.Flags().GetString("button-color")
		body["button_color"] = v
	}
	if cmd.Flags().Changed("button-corners") {
		v, _ := cmd.Flags().GetInt("button-corners")
		body["button_corners"] = v
	}
	if cmd.Flags().Changed("button-text") {
		v, _ := cmd.Flags().GetString("button-text")
		body["button_text"] = v
	}
	if cmd.Flags().Changed("font-family") {
		v, _ := cmd.Flags().GetString("font-family")
		body["font_family"] = v
	}
	if cmd.Flags().Changed("heading-font-family") {
		v, _ := cmd.Flags().GetString("heading-font-family")
		body["heading_font_family"] = v
	}
	if cmd.Flags().Changed("keywords") {
		v, _ := cmd.Flags().GetStringSlice("keywords")
		body["keywords"] = v
	}
	if cmd.Flags().Changed("links-color") {
		v, _ := cmd.Flags().GetString("links-color")
		body["links_color"] = v
	}
	if cmd.Flags().Changed("logo-url") {
		v, _ := cmd.Flags().GetString("logo-url")
		body["logo_url"] = v
	}
	if cmd.Flags().Changed("privacy-policy-url") {
		v, _ := cmd.Flags().GetString("privacy-policy-url")
		body["privacy_policy_url"] = v
	}
	if cmd.Flags().Changed("selling-points") {
		v, _ := cmd.Flags().GetStringSlice("selling-points")
		body["selling_points"] = v
	}
	if cmd.Flags().Changed("sender-email") {
		v, _ := cmd.Flags().GetString("sender-email")
		body["sender_email"] = v
	}
	if cmd.Flags().Changed("sender-name") {
		v, _ := cmd.Flags().GetString("sender-name")
		body["sender_name"] = v
	}
	if cmd.Flags().Changed("social-links") {
		v, _ := cmd.Flags().GetStringSlice("social-links")
		body["social_links"] = v
	}
	if cmd.Flags().Changed("socials") {
		v, _ := cmd.Flags().GetString("socials")
		body["socials"] = utils.JSONValue(v, "socials")
	}
	if cmd.Flags().Changed("text-color") {
		v, _ := cmd.Flags().GetString("text-color")
		body["text_color"] = v
	}
	if cmd.Flags().Changed("topics") {
		v, _ := cmd.Flags().GetStringSlice("topics")
		body["topics"] = v
	}
	if cmd.Flags().Changed("website-url") {
		v, _ := cmd.Flags().GetString("website-url")
		body["website_url"] = v
	}
	return body
}
