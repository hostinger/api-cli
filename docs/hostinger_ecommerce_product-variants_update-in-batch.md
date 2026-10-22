## hostinger ecommerce product-variants update-in-batch

Update product variants in batch

### Synopsis

Update up to 100 existing variants in place by id — title, SKU, inventory, stock tracking
and prices. Variants omitted from the request are left untouched. Prices replace the variant's
existing prices in full. Returns the updated variants.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --variants
    [
      {
        "inventory_quantity": 42,
        "manage_inventory": true,
        "prices": [
          {
            "amount": 1999,
            "currency": "usd",
            "sale_amount": 1499
          }
        ],
        "sku": "TSHIRT-RED-L",
        "title": "Red / L",
        "variant_id": "variant_01J8Z5F8W9K8M4A7B3C2D1E0FG"
      }
    ]

```
hostinger ecommerce product-variants update-in-batch <store_id> <product_id> [flags]
```

### Examples

```
  hostinger ecommerce product-variants update-in-batch store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG --variants @variants.json
```

### Options

```
  -h, --help              help for update-in-batch
      --variants string   Variants to update in place by id, up to 100. Variants omitted from the list are left untouched. (JSON, @file or -; shape in the help above)
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

