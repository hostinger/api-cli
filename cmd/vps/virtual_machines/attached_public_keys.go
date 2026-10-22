package virtual_machines

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AttachedPublicKeysCmd = &cobra.Command{
	Use:     "attached-public-keys <virtual-machine-id>",
	Short:   "Get attached public keys",
	Long:    "Deprecated: use `GET /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys` instead.\n\nRetrieve public keys attached to a specified virtual machine.\n\nUse this endpoint to view SSH keys configured for specific VPS instances.",
	Example: "  hostinger vps virtual-machines attached-public-keys 1268054",
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
		params := attachedPublicKeysParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSGetAttachedPublicKeysV1WithResponse(ctx, utils.StringToInt(args[0]), params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(AttachedPublicKeysCmd, utils.IntArg("virtual-machine-id"))
	AttachedPublicKeysCmd.Flags().IntP("page", "", 0, "Page number")
	AttachedPublicKeysCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	AttachedPublicKeysCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	AttachedPublicKeysCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func attachedPublicKeysParams(cmd *cobra.Command) *client.VPSGetAttachedPublicKeysV1Params {
	params := &client.VPSGetAttachedPublicKeysV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
