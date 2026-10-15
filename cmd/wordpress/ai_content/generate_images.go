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
	Long:  "Start generating one or more images from a prompt and an image type for the specified\nWordPress installation. The generation runs asynchronously: the response is returned\nright away with `status: in-progress` and a `generation_id`. Poll\nGET /api/hosting/v1/accounts/{username}/wordpress/{software}/content/generations/{generation_id}\nevery few seconds until `status` is `completed` or `failed`.\n\nA completed generation returns temporary image URLs together with the time they expire;\ndownload the images before then. Images match the requested aspect ratio and are never\nsmaller than the requested size; logos are transparent images cropped to the mark.\n\nOnly one image generation can run per installation at a time: while one is still in\nprogress the endpoint returns 409 without the ID of the running generation, so a caller\nthat did not receive the `202` has to wait for it to finish (up to a few minutes) before\nstarting a new one. The endpoint is limited to 10 requests per minute.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
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
