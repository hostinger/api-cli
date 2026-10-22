package virtual_machines

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetHostnameCmd = &cobra.Command{
	Use:     "set-hostname <virtual-machine-id>",
	Short:   "Set hostname",
	Long:    "Set hostname for a specified virtual machine.\n\nChanging hostname does not update PTR record automatically.\nIf you want your virtual machine to be reachable by a hostname,\nyou need to point your domain A/AAAA records to virtual machine IP as well.\n\nUse this endpoint to configure custom hostnames for VPS instances.",
	Example: "  hostinger vps virtual-machines set-hostname 1268054 --hostname my.server.tld",
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
		body, err := setHostnameBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSSetHostnameV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(SetHostnameCmd, utils.IntArg("virtual-machine-id"))
	SetHostnameCmd.Flags().StringP("hostname", "", "", "")
	SetHostnameCmd.MarkFlagRequired("hostname")
	SetHostnameCmd.RegisterFlagCompletionFunc("hostname", cobra.NoFileCompletions)
	utils.RegisterWaitFlags(SetHostnameCmd)
}

func setHostnameBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	hostnameVal, _ := cmd.Flags().GetString("hostname")
	body["hostname"] = hostnameVal
	return body, nil
}
