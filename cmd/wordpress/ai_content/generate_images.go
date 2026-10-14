package ai_content

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

var GenerateImagesCmd = &cobra.Command{
	Use:   "generate-images <username> <software>",
	Short: "Generate images",
	Long:  "Generate one or more images from a prompt and an image type for the specified\nWordPress installation. Images are returned inline as base64-encoded PNGs. They match\nthe requested aspect ratio and are never smaller than the requested size; logos are\ntransparent PNGs cropped to the mark.\n\nGeneration is synchronous and can take up to a minute. The endpoint is limited to\n10 requests per minute.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "image-type", []string{"hero", "background", "inline", "gallery", "product", "logo"})
		payload, err := json.Marshal(generateImagesBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingGenerateImagesV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	GenerateImagesCmd.Flags().StringP("color-guidance", "", "", "Colour direction for the logo type only. It is not rendered as text.")
	GenerateImagesCmd.Flags().IntP("height", "", 0, "Output height in pixels. Defaults per image type and is ignored for the logo type.")
	GenerateImagesCmd.Flags().StringP("image-type", "", "", "Kind of image to generate. It sets the default size and cropping. (one of: hero, background, inline, gallery, product, logo)")
	GenerateImagesCmd.Flags().IntP("num-of-images", "", 1, "How many images to generate")
	GenerateImagesCmd.Flags().StringP("prompt", "", "", "What the image should show. For the logo type, the wordmark and mark to render.")
	GenerateImagesCmd.Flags().StringP("style", "", "", "Style tokens appended to photographic prompts. Ignored for the logo type.")
	GenerateImagesCmd.Flags().IntP("width", "", 0, "Output width in pixels. Defaults per image type and is ignored for the logo type.")
	GenerateImagesCmd.MarkFlagRequired("image-type")
	GenerateImagesCmd.MarkFlagRequired("prompt")
}

func generateImagesBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("color-guidance") {
		v, _ := cmd.Flags().GetString("color-guidance")
		body["color_guidance"] = v
	}
	if cmd.Flags().Changed("height") {
		v, _ := cmd.Flags().GetInt("height")
		body["height"] = v
	}
	imageTypeVal, _ := cmd.Flags().GetString("image-type")
	body["image_type"] = imageTypeVal
	if cmd.Flags().Changed("num-of-images") {
		v, _ := cmd.Flags().GetInt("num-of-images")
		body["num_of_images"] = v
	}
	promptVal, _ := cmd.Flags().GetString("prompt")
	body["prompt"] = promptVal
	if cmd.Flags().Changed("style") {
		v, _ := cmd.Flags().GetString("style")
		body["style"] = v
	}
	if cmd.Flags().Changed("width") {
		v, _ := cmd.Flags().GetInt("width")
		body["width"] = v
	}
	return body
}
