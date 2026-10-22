## hostinger ecommerce products upload-and-attach-image

Upload and attach a product image

### Synopsis

Fetch a raster image (JPEG, PNG, GIF or WebP, max 15MB) from a URL and attach it to a product in a
single call. Image downloads require HTTPS on port 443 without embedded credentials. At most one redirect
is allowed, and its destination must meet the same requirements. Private or reserved network
destinations, unsupported URLs and longer redirect chains are rejected. The image is virus-scanned
and validated by content, then stored on the CDN. Set is_thumbnail to make it the product's primary image.

```
hostinger ecommerce products upload-and-attach-image <store_id> <product_id> [flags]
```

### Examples

```
  hostinger ecommerce products upload-and-attach-image store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help                 help for upload-and-attach-image
      --image-url string     Publicly reachable URL of a raster image (JPEG, PNG, GIF or WebP), maximum 15MB. Fetching the image requires HTTPS on port 443 without embedded credentials. At most one redirect is allowed; its destination must meet the same URL requirements. Private or reserved network destinations, unsupported URLs and longer redirect chains are rejected. The image is fetched, virus-scanned and validated by content, then stored on the CDN. SVG is not accepted. Provide either this or object_name.
      --is-thumbnail         When true, the image becomes the product's thumbnail (primary image). When omitted, it becomes the thumbnail only if the product does not have one yet.
      --object-name string   Key returned by the upload-url endpoint. Provide this instead of image_url to attach an uploaded image.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

