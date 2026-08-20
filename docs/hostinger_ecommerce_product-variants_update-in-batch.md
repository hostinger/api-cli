## hostinger ecommerce product-variants update-in-batch

Update product variants in batch

### Synopsis

Update up to 100 existing variants in place by id — title, inventory, stock tracking and
prices. Variants omitted from the request are left untouched. Prices replace the variant's
existing prices in full. Returns the updated variants.

```
hostinger ecommerce product-variants update-in-batch <store_id> <product_id> [flags]
```

### Options

```
  -h, --help              help for update-in-batch
      --variants string   Variants to update in place by id, up to 100. Variants omitted from the list are left untouched. (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce product-variants](hostinger_ecommerce_product-variants.md)	 - Product variants commands

