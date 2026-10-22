package public_keys

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AttachCmd = &cobra.Command{
	Use:     "attach <virtual-machine-id>",
	Short:   "Attach public key",
	Long:    "Deprecated: use `POST /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys` instead.\n\nAttach existing public keys from your account to a specified virtual machine.\n\nMultiple keys can be attached to a single virtual machine.\n\nUse this endpoint to enable SSH key authentication for VPS instances.",
	Example: "  hostinger vps public-keys attach 1268054 --ids 18232,10230230",
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
		body, err := attachBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSAttachPublicKeyV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(AttachCmd, utils.IntArg("virtual-machine-id"))
	AttachCmd.Flags().IntSliceP("ids", "", nil, "Public Key IDs to attach")
	AttachCmd.MarkFlagRequired("ids")
	AttachCmd.RegisterFlagCompletionFunc("ids", cobra.NoFileCompletions)
	utils.RegisterWaitFlags(AttachCmd)
}

func attachBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	idsVal, _ := cmd.Flags().GetIntSlice("ids")
	body["ids"] = idsVal
	return body, nil
}
