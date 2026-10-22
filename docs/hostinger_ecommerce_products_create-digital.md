## hostinger ecommerce products create-digital

Create digital product

### Synopsis

Create a published digital product with a single variant and an optional external download link.

```
hostinger ecommerce products create-digital <store_id> [flags]
```

### Examples

```
  hostinger ecommerce products create-digital store_01J8Z5F8W9K8M4A7B3C2D1E0FG --name 'My eBook' --price 999
```

### Options

```
      --currency string       ISO 4217 currency code. Defaults to the store's default currency when omitted.
      --description string    The product description.
      --download-url string   Optional external download link delivered to the customer after purchase.
  -h, --help                  help for create-digital
      --name string           The product name.
      --price int             Price in the smallest currency unit (e.g. cents). Must be positive.
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

