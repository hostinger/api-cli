## hostinger wordpress ai-content generate-images

Generate images

### Synopsis

Generate one or more images from a prompt and an image type for the specified
WordPress installation. Images are returned inline as base64-encoded PNGs. They match
the requested aspect ratio and are never smaller than the requested size; logos are
transparent PNGs cropped to the mark.

Generation is synchronous and can take up to a minute. The endpoint is limited to
10 requests per minute.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress ai-content generate-images <username> <software> [flags]
```

### Options

```
      --color-guidance string   Colour direction for the logo type only. It is not rendered as text.
      --height int              Output height in pixels. Defaults per image type and is ignored for the logo type.
  -h, --help                    help for generate-images
      --image-type string       Kind of image to generate. It sets the default size and cropping. (one of: hero, background, inline, gallery, product, logo)
      --num-of-images int       How many images to generate (default 1)
      --prompt string           What the image should show. For the logo type, the wordmark and mark to render.
      --style string            Style tokens appended to photographic prompts. Ignored for the logo type.
      --width int               Output width in pixels. Defaults per image type and is ignored for the logo type.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress ai-content](hostinger_wordpress_ai-content.md)	 - AI Content commands

