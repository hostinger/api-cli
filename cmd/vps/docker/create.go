package docker

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

var CreateCmd = &cobra.Command{
	Use:   "create <virtual-machine-id>",
	Short: "Create new project",
	Long:  "Deploy new project from docker-compose.yaml contents or download contents from URL. \n\nURL can be Github repository url in format https://github.com/[user]/[repo]\nand it will be automatically resolved to docker-compose.yaml file in\nmaster branch. Any other URL provided must return docker-compose.yaml\nfile contents.\n\nIf project with the same name already exists, existing project will be replaced.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSCreateNewProjectV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("content", "", "", "URL pointing to docker-compose.yaml file, Github repository or raw YAML content of the compose file")
	CreateCmd.Flags().StringP("environment", "", "", "Project environment variables")
	CreateCmd.Flags().StringP("project-name", "", "", "Docker Compose project name using alphanumeric characters, dashes, and underscores only")
	CreateCmd.MarkFlagRequired("content")
	CreateCmd.MarkFlagRequired("project-name")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	if cmd.Flags().Changed("environment") {
		v, _ := cmd.Flags().GetString("environment")
		body["environment"] = v
	}
	projectNameVal, _ := cmd.Flags().GetString("project-name")
	body["project_name"] = projectNameVal
	return body
}
