## hostinger wordpress ai-content generate-images

Generate images

### Synopsis

Start generating one or more images from a prompt and an image type for the specified
WordPress installation. The generation runs asynchronously: the response is returned
right away with `status: in-progress` and a `generation_id`. Poll
GET /api/hosting/v1/accounts/{username}/wordpress/{software}/content/generations/{generation_id}
every few seconds until `status` is `completed` or `failed`.

A completed generation returns temporary image URLs together with the time they expire;
download the images before then. Images match the requested aspect ratio and are never
smaller than the requested size; logos are transparent images cropped to the mark.

Only one image generation can run per installation at a time: while one is still in
progress the endpoint returns 409 without the ID of the running generation, so a caller
that did not receive the `202` has to wait for it to finish (up to a few minutes) before
starting a new one. The endpoint is limited to 10 requests per minute.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

The hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger wordpress ai-content generate-images [username] <software> [flags]
```

### Examples

```
  hostinger wordpress ai-content generate-images u123456789 1232456789 --image-type hero --prompt 'A cozy artisan bakery interior at sunrise with fresh bread on wooden shelves'
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress ai-content](hostinger_wordpress_ai-content.md)	 - AI Content commands

