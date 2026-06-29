## hostinger ecommerce products create-physical

Create physical product

### Synopsis

Create a published physical product with a single variant priced in the store currency.

```
hostinger ecommerce products create-physical <store_id> [flags]
```

### Options

```
      --currency string      ISO 4217 currency code. Defaults to the store's default currency when omitted.
      --description string   The product description.
  -h, --help                 help for create-physical
      --name string          The product name.
      --price int            Price in the smallest currency unit (e.g. cents). Must be positive.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

