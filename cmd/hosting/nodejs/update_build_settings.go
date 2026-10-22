package nodejs

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateBuildSettingsCmd = &cobra.Command{
	Use:     "update-build-settings <username> <domain>",
	Short:   "Update Node.js build settings",
	Long:    "Replaces the build settings stored for the website. Send the full set: `node_version` is\nrequired and every nullable field you omit is stored as null. Creates the settings when\nnone exist yet.\n\nThis does not start a build. Stored settings drive Git auto-deployment builds; a build\nstarted through the API uses the values sent in that request, so to rebuild with corrected\nsettings call `Start Node.js build` with the same values. Typical fixes: a wrong `app_type`\nafter auto-detection, or a missing `entry_file` for express, fastify, nest, nuxt and hono\napps.",
	Example: "  hostinger hosting nodejs update-build-settings u123456789 mydomain.tld --node-version 20",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "app-type", []string{"create-react-app", "gatsby", "vite", "angular", "react", "vue", "parcel", "next", "nuxt", "nest", "express", "fastify", "astro", "svelte", "svelte-kit", "hono", "react-router", "nitro", "other"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "package-manager", []string{"npm", "yarn", "pnpm"}); err != nil {
			return err
		}
		body, err := updateBuildSettingsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdateNodeJsBuildSettingsV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateBuildSettingsCmd, utils.StringArg("username"), utils.StringArg("domain"))
	UpdateBuildSettingsCmd.Flags().StringP("app-type", "", "", "Node.js application framework. Set it explicitly when auto-detection picked the wrong one. (one of: create-react-app, gatsby, vite, angular, react, vue, parcel, next, nuxt, nest, express, fastify, astro, svelte, svelte-kit, hono, react-router, nitro, other)")
	UpdateBuildSettingsCmd.Flags().StringP("build-script", "", "", "The package.json script that builds the application")
	UpdateBuildSettingsCmd.Flags().StringP("entry-file", "", "", "The main entry point file for the application (required for express, fastify, nest, nuxt and hono app types)")
	UpdateBuildSettingsCmd.Flags().IntP("node-version", "", 0, "Node.js major version (one of: 18, 20, 22, 24)")
	UpdateBuildSettingsCmd.Flags().StringP("output-directory", "", "", "Build output directory relative to the root directory")
	UpdateBuildSettingsCmd.Flags().StringP("package-manager", "", "", "Package manager used to install dependencies (one of: npm, yarn, pnpm)")
	UpdateBuildSettingsCmd.Flags().StringP("root-directory", "", "", "Application root directory (where package.json is located) relative to public_html. Omit it, or send \".\", for public_html itself.")
	UpdateBuildSettingsCmd.MarkFlagRequired("node-version")
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("app-type", cobra.FixedCompletions([]cobra.Completion{"create-react-app", "gatsby", "vite", "angular", "react", "vue", "parcel", "next", "nuxt", "nest", "express", "fastify", "astro", "svelte", "svelte-kit", "hono", "react-router", "nitro", "other"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("build-script", cobra.NoFileCompletions)
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("entry-file", cobra.NoFileCompletions)
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("node-version", cobra.FixedCompletions([]cobra.Completion{"18", "20", "22", "24"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("output-directory", cobra.NoFileCompletions)
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("package-manager", cobra.FixedCompletions([]cobra.Completion{"npm", "yarn", "pnpm"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateBuildSettingsCmd.RegisterFlagCompletionFunc("root-directory", cobra.NoFileCompletions)
}

func updateBuildSettingsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("app-type") {
		v, _ := cmd.Flags().GetString("app-type")
		body["app_type"] = v
	}
	if cmd.Flags().Changed("build-script") {
		v, _ := cmd.Flags().GetString("build-script")
		body["build_script"] = v
	}
	if cmd.Flags().Changed("entry-file") {
		v, _ := cmd.Flags().GetString("entry-file")
		body["entry_file"] = v
	}
	nodeVersionVal, _ := cmd.Flags().GetInt("node-version")
	body["node_version"] = nodeVersionVal
	if cmd.Flags().Changed("output-directory") {
		v, _ := cmd.Flags().GetString("output-directory")
		body["output_directory"] = v
	}
	if cmd.Flags().Changed("package-manager") {
		v, _ := cmd.Flags().GetString("package-manager")
		body["package_manager"] = v
	}
	if cmd.Flags().Changed("root-directory") {
		v, _ := cmd.Flags().GetString("root-directory")
		body["root_directory"] = v
	}
	return body, nil
}
