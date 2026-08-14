package products

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UploadAndAttachImageCmd = &cobra.Command{
	Use:   "upload-and-attach-image <store_id> <product_id>",
	Short: "Upload and attach a product image",
	Long:  "Fetch a raster image (JPEG, PNG, GIF or WebP, max 15MB) from a URL and attach it to a product in a\nsingle call. The image is virus-scanned and validated by content, then stored on the CDN. Set\nis_thumbnail to make it the product's primary image.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(uploadAndAttachImageBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceUploadAndAttachAProductImageV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UploadAndAttachImageCmd.Flags().StringP("image-url", "", "", "Publicly reachable URL of the raster image (JPEG, PNG, GIF or WebP), maximum 15MB. The image is\nfetched, virus-scanned and validated by content, then stored on the CDN. SVG is not accepted.\nProvide either this or object_name.")
	UploadAndAttachImageCmd.Flags().BoolP("is-thumbnail", "", false, "When true, the image becomes the product's thumbnail (primary image). When omitted, it becomes the\nthumbnail only if the product does not have one yet.")
	UploadAndAttachImageCmd.Flags().StringP("object-name", "", "", "Key returned by the upload-url endpoint. Provide this instead of image_url to attach an uploaded image.")
}

func uploadAndAttachImageBody(cmd *cobra.Command) map[string]any {
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
	return body
}
