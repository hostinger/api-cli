## hostinger ecommerce product-variants list

List product variants

### Synopsis

List a product's variants, ordered by rank, with their options, prices and inventory.
Prices are integers in the smallest currency unit and live on variants.

```
hostinger ecommerce product-variants list <store_id> <product_id> [flags]
```

### Examples

```
  hostinger ecommerce product-variants list store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --all        Fetch every page and print all items
  -h, --help       help for list
      --page int   Page number
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

