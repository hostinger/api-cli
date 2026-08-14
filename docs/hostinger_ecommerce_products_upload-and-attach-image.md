## hostinger ecommerce products upload-and-attach-image

Upload and attach a product image

### Synopsis

Fetch a raster image (JPEG, PNG, GIF or WebP, max 15MB) from a URL and attach it to a product in a
single call. The image is virus-scanned and validated by content, then stored on the CDN. Set
is_thumbnail to make it the product's primary image.

```
hostinger ecommerce products upload-and-attach-image <store_id> <product_id> [flags]
```

### Options

```
  -h, --help               help for upload-and-attach-image
      --image-url string   Publicly reachable URL of the raster image (JPEG, PNG, GIF or WebP), maximum 15MB. The image is
                           fetched, virus-scanned and validated by content, then stored on the CDN. SVG is not accepted.
      --is-thumbnail       When true, the image becomes the product's thumbnail (primary image). When omitted, it becomes the
                           thumbnail only if the product does not have one yet.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

