## hapi ecommerce products create-digital

Create digital product

### Synopsis

Create a published digital product with a single variant and an optional external download link.

```
hapi ecommerce products create-digital <store_id> [flags]
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
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi ecommerce products](hapi_ecommerce_products.md)	 - Products commands

