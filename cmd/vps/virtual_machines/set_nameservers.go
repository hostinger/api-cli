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

var SetNameserversCmd = &cobra.Command{
	Use:     "set-nameservers <virtual-machine-id>",
	Short:   "Set nameservers",
	Long:    "Set nameservers for a specified virtual machine.\n\nBe aware, that improper nameserver configuration can lead to the virtual\nmachine being unable to resolve domain names.\n\nUse this endpoint to configure custom DNS resolvers for VPS instances.",
	Example: "  hostinger vps virtual-machines set-nameservers 1268054 --ns1 4.3.2.1",
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
		body, err := setNameserversBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSSetNameserversV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(SetNameserversCmd, utils.IntArg("virtual-machine-id"))
	SetNameserversCmd.Flags().StringP("ns1", "", "", "")
	SetNameserversCmd.Flags().StringP("ns2", "", "", "")
	SetNameserversCmd.Flags().StringP("ns3", "", "", "")
	SetNameserversCmd.MarkFlagRequired("ns1")
	SetNameserversCmd.RegisterFlagCompletionFunc("ns1", cobra.NoFileCompletions)
	SetNameserversCmd.RegisterFlagCompletionFunc("ns2", cobra.NoFileCompletions)
	SetNameserversCmd.RegisterFlagCompletionFunc("ns3", cobra.NoFileCompletions)
	utils.RegisterWaitFlags(SetNameserversCmd)
}

func setNameserversBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	ns1Val, _ := cmd.Flags().GetString("ns1")
	body["ns1"] = ns1Val
	if cmd.Flags().Changed("ns2") {
		v, _ := cmd.Flags().GetString("ns2")
		body["ns2"] = v
	}
	if cmd.Flags().Changed("ns3") {
		v, _ := cmd.Flags().GetString("ns3")
		body["ns3"] = v
	}
	return body, nil
}
