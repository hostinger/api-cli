package recovery

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

var StartCmd = &cobra.Command{
	Use:   "start <virtual-machine-id>",
	Short: "Start recovery mode",
	Long:  "Initiate recovery mode for a specified virtual machine.\n\nRecovery mode is a special state that allows users to perform system rescue operations, \nsuch as repairing file systems, recovering data, or troubleshooting issues that prevent the virtual machine \nfrom booting normally. \n\nVirtual machine will boot recovery disk image and original disk image will be mounted in `/mnt` directory.\n\nUse this endpoint to enable system rescue operations on VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(startBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSStartRecoveryModeV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	StartCmd.Flags().StringP("root-password", "", "", "Temporary root password for recovery mode")
	StartCmd.MarkFlagRequired("root-password")
}

func startBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	rootPasswordVal, _ := cmd.Flags().GetString("root-password")
	body["root_password"] = rootPasswordVal
	return body
}
