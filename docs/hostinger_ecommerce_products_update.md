## hostinger ecommerce products update

Update a product

### Synopsis

Update a product's name, description or status. Set status to published to make it buyable,
draft to hide it, or archived to retire it. Variants, prices and inventory are managed
through the variant endpoints, not here. Returns the updated product summary.

```
hostinger ecommerce products update <store_id> <product_id> [flags]
```

### Options

```
      --description string   The product description.
  -h, --help                 help for update
      --name string          The product name.
      --status string        Set "published" to make the product buyable, "draft" to hide it, or "archived" to retire it. (one of: draft, published, archived)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce products](hostinger_ecommerce_products.md)	 - Products commands

