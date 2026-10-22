package products

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UploadAndAttachImageCmd = &cobra.Command{
	Use:               "upload-and-attach-image <store_id> <product_id>",
	Short:             "Upload and attach a product image",
	Long:              "Fetch a raster image (JPEG, PNG, GIF or WebP, max 15MB) from a URL and attach it to a product in a\nsingle call. Image downloads require HTTPS on port 443 without embedded credentials. At most one redirect\nis allowed, and its destination must meet the same requirements. Private or reserved network\ndestinations, unsupported URLs and longer redirect chains are rejected. The image is virus-scanned\nand validated by content, then stored on the CDN. Set is_thumbnail to make it the product's primary image.",
	Example:           "  hostinger ecommerce products upload-and-attach-image store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := uploadAndAttachImageBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceUploadAndAttachAProductImageV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UploadAndAttachImageCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	UploadAndAttachImageCmd.Flags().StringP("image-url", "", "", "Publicly reachable URL of a raster image (JPEG, PNG, GIF or WebP), maximum 15MB. Fetching the image requires HTTPS on port 443 without embedded credentials. At most one redirect is allowed; its destination must meet the same URL requirements. Private or reserved network destinations, unsupported URLs and longer redirect chains are rejected. The image is fetched, virus-scanned and validated by content, then stored on the CDN. SVG is not accepted. Provide either this or object_name.")
	UploadAndAttachImageCmd.Flags().BoolP("is-thumbnail", "", false, "When true, the image becomes the product's thumbnail (primary image). When omitted, it becomes the thumbnail only if the product does not have one yet.")
	UploadAndAttachImageCmd.Flags().StringP("object-name", "", "", "Key returned by the upload-url endpoint. Provide this instead of image_url to attach an uploaded image.")
	UploadAndAttachImageCmd.RegisterFlagCompletionFunc("image-url", cobra.NoFileCompletions)
	UploadAndAttachImageCmd.RegisterFlagCompletionFunc("is-thumbnail", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UploadAndAttachImageCmd.RegisterFlagCompletionFunc("object-name", cobra.NoFileCompletions)
}

func uploadAndAttachImageBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("image-url") {
		v, _ := cmd.Flags().GetString("image-url")
		body["image_url"] = v
	}
	if cmd.Flags().Changed("is-thumbnail") {
		v, _ := cmd.Flags().GetBool("is-thumbnail")
		body["is_thumbnail"] = v
	}
	if cmd.Flags().Changed("object-name") {
		v, _ := cmd.Flags().GetString("object-name")
		body["object_name"] = v
	}
	return body, nil
}
