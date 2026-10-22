package docker

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:     "create <virtual-machine-id>",
	Short:   "Create new project",
	Long:    "Deploy new project from docker-compose.yaml contents or download contents from URL.\n\nURL can be Github repository url in format https://github.com/[user]/[repo]\nand it will be automatically resolved to docker-compose.yaml file in\nmaster branch. Any other URL provided must return docker-compose.yaml\nfile contents.\n\nIf project with the same name already exists, existing project will be replaced.",
	Example: "  hostinger vps docker create 1268054 --content value --project-name my-project-1",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetVirtualMachinesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSCreateNewProjectV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		var waitErr error
		if wait, _ := cmd.Flags().GetBool("wait"); wait && status < 300 {
			respBody, status, waitErr = utils.Wait(cmd, respBody, utils.WaitSpec{
				IDField:    "id",
				StateField: "state",
				Done:       []string{"success"},
				Failed:     []string{"error"},
			}, func(ctx context.Context, id string) ([]byte, int, error) {
				r, err := api.Request().VPSGetActionDetailsV1WithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(id))
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			})
			if respBody == nil {
				return waitErr
			}
		}
		if err := output.Format(cmd, respBody, status); err != nil {
			return err
		}
		return waitErr
	},
}

func init() {
	utils.RegisterArgs(CreateCmd, utils.IntArg("virtual-machine-id"))
	CreateCmd.Flags().StringP("content", "", "", "URL pointing to docker-compose.yaml file, Github repository or raw YAML content of the compose file")
	CreateCmd.Flags().StringP("environment", "", "", "Project environment variables")
	CreateCmd.Flags().StringP("project-name", "", "", "Docker Compose project name using alphanumeric characters, dashes, and underscores only")
	CreateCmd.MarkFlagRequired("content")
	CreateCmd.MarkFlagRequired("project-name")
	CreateCmd.RegisterFlagCompletionFunc("content", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("environment", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("project-name", cobra.NoFileCompletions)
	utils.RegisterWaitFlags(CreateCmd)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	if cmd.Flags().Changed("environment") {
		v, _ := cmd.Flags().GetString("environment")
		body["environment"] = v
	}
	projectNameVal, _ := cmd.Flags().GetString("project-name")
	body["project_name"] = projectNameVal
	return body, nil
}
