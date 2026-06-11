package ptr

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var CreateCmd = &cobra.Command{
	Use:   "create <virtual machine ID> <IP address ID>",
	Short: "Create PTR record",
	Long:  `This endpoint creates or updates a PTR (Pointer) record for a specified IP address of a virtual machine.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSCreatePTRRecordV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]), ptrCreateRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("domain", "", "", "Pointer record domain")

	CreateCmd.MarkFlagRequired("domain")
}

func ptrCreateRequest(cmd *cobra.Command) client.VPSCreatePTRRecordV1JSONRequestBody {
	domain, _ := cmd.Flags().GetString("domain")

	return client.VPSCreatePTRRecordV1JSONRequestBody{
		Domain: domain,
	}
}
