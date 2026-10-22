## hostinger ecommerce product-variants create

Create a product variant

### Synopsis

Add a variant to a product along one or more option dimensions (e.g. Size, Color). Options
missing from the product are created automatically; provide a value for every option the
product already has. Prices are integers in the smallest currency unit and default to the
store currency. Returns the created variant.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --options
    [
      {
        "name": "Size",
        "value": "M"
      }
    ]

  --prices
    [
      {
        "amount": 1999,
        "currency": "usd",
        "sale_amount": 1499
      }
    ]

```
hostinger ecommerce product-variants create <store_id> <product_id> [flags]
```

### Examples

```
  hostinger ecommerce product-variants create store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG --options '[{"name":"Size","value":"M"}]'
```

### Options

```
  -h, --help                     help for create
      --inventory-quantity int   Units in stock. Defaults to 0.
      --manage-inventory         Whether stock is tracked for this variant. Defaults to false.
      --options string           Option name/value pairs that distinguish this variant, e.g. [{name: Size, value: M}]. Options missing from the product are created; provide a value for every option the product already has. (JSON, @file or -; shape in the help above)
      --prices string            Prices per currency. Amounts are integers in the smallest currency unit. A free item is amount: 0. (JSON, @file or -; shape in the help above)
      --sku string               The variant SKU.
      --title string             The variant title. Defaults to the option values joined with ' / ' (e.g. 'Red / L').
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

