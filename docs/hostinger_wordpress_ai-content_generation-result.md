## hostinger wordpress ai-content generation-result

Get generation result

### Synopsis

Get the result of an AI content generation started for the specified WordPress
installation, such as an image generation started with
POST /api/hosting/v1/accounts/{username}/wordpress/{software}/content/images/generate.

While the generation is still running only `status: in-progress` is returned; poll every
few seconds. Once `completed`, `result` holds the generated content shaped according to
`type`: temporary image URLs together with the time they expire for `image-generation`, or
the generated website document for `website-generation` (started from the Hostinger
WordPress plugin). A `failed` generation has no `result`. Unknown or expired generation
IDs return 404.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress ai-content generation-result <username> <software> <generation_id> [flags]
```

### Options

```
  -h, --help          help for generation-result
      --type string   Generation type to look up. When omitted, every generation type is checked. (one of: website-generation, image-generation)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress ai-content](hostinger_wordpress_ai-content.md)	 - AI Content commands

