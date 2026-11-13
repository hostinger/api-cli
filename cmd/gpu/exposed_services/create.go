package exposed_services

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
	Use:     "create <instance-id>",
	Short:   "Create exposed service",
	Long:    "Expose a service running on a specified GPU instance to the internet.\n\nThe service is exposed asynchronously: a successful request returns immediately, and the service\nappears in the exposed services list once it gets its public hostname and external port.\nFails when the instance already has the maximum of 4 exposed services.\n\nUse this endpoint to make an application on your GPU instance publicly reachable.",
	Example: "  hostinger gpu exposed-services create 6325 --internal-port 8000 --name Ollama --protocol tcp --type http",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().GPUGetInstancesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "protocol", []string{"tcp", "udp"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "type", []string{"http", "https"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUCreateExposedServiceV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.IntArg("instance-id"))
	CreateCmd.Flags().IntP("internal-port", "", 0, "Port the service listens on inside the instance, unique within the instance")
	CreateCmd.Flags().StringP("name", "", "", "Service name, unique within the instance")
	CreateCmd.Flags().StringP("protocol", "", "", "(one of: tcp, udp)")
	CreateCmd.Flags().StringP("type", "", "", "(one of: http, https)")
	CreateCmd.MarkFlagRequired("internal-port")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("protocol")
	CreateCmd.MarkFlagRequired("type")
	CreateCmd.RegisterFlagCompletionFunc("internal-port", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("protocol", cobra.FixedCompletions([]cobra.Completion{"tcp", "udp"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"http", "https"}, cobra.ShellCompDirectiveNoFileComp))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	internalPortVal, _ := cmd.Flags().GetInt("internal-port")
	body["internal_port"] = internalPortVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	protocolVal, _ := cmd.Flags().GetString("protocol")
	body["protocol"] = protocolVal
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	return body, nil
}
