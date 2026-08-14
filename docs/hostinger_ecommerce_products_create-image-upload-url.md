## hostinger ecommerce products create-image-upload-url

Create a product image upload URL

### Synopsis

Returns a signed URL to upload a product image to (multipart/form-data POST). Then call the
attach-image endpoint with the returned object_name to scan and attach it to the product.

```
hostinger ecommerce products create-image-upload-url <store_id> <product_id> [flags]
```

### Options

```
  -h, --help   help for create-image-upload-url
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

