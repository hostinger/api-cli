## hostinger ecommerce product-variants delete

Delete a product variant

### Synopsis

Delete a single variant from the product.

```
hostinger ecommerce product-variants delete <store_id> <product_id> <variant_id> [flags]
```

### Examples

```
  hostinger ecommerce product-variants delete store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG variant_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

* [hostinger ecommerce product-variants](hostinger_ecommerce_product-variants.md)	 - Product variants commands

