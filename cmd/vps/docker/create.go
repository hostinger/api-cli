package docker

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
	Use:   "create <virtual machine ID>",
	Short: "Create Docker project",
	Long:  `This endpoint creates and starts a new Docker Compose project on a specified virtual machine.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSCreateNewProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), projectCreateRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("project-name", "", "", "Docker Compose project name (alphanumeric characters, dashes and underscores only)")
	CreateCmd.Flags().StringP("content", "", "", "URL to a docker-compose.yaml file, a GitHub repository, or raw YAML content of the compose file")
	CreateCmd.Flags().StringP("environment", "", "", "Project environment variables")

	CreateCmd.MarkFlagRequired("project-name")
	CreateCmd.MarkFlagRequired("content")
}

func projectCreateRequest(cmd *cobra.Command) client.VPSCreateNewProjectV1JSONRequestBody {
	projectName, _ := cmd.Flags().GetString("project-name")
	content, _ := cmd.Flags().GetString("content")
	environment, _ := cmd.Flags().GetString("environment")

	return client.VPSCreateNewProjectV1JSONRequestBody{
		ProjectName: projectName,
		Content:     content,
		Environment: utils.StringPtrOrNil(environment),
	}
}
