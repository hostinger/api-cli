## hostinger ecommerce product-variants create

Create a product variant

### Synopsis

Add a variant to a product along one or more option dimensions (e.g. Size, Color). Options
missing from the product are created automatically; provide a value for every option the
product already has. Prices are integers in the smallest currency unit and default to the
store currency. Returns the created variant.

```
hostinger ecommerce product-variants create <store_id> <product_id> [flags]
```

### Options

```
  -h, --help                     help for create
      --inventory-quantity int   Units in stock. Defaults to 0.
      --manage-inventory         Whether stock is tracked for this variant. Defaults to false.
      --options string           Option name/value pairs that distinguish this variant, e.g. [{name: Size, value: M}]. Options missing from the product are created; provide a value for every option the product already has. (JSON)
      --prices string            Prices per currency. Amounts are integers in the smallest currency unit. A free item is amount: 0. (JSON)
      --sku string               The variant SKU.
      --title string             The variant title. Defaults to the option values joined with ' / ' (e.g. 'Red / L').
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce product-variants](hostinger_ecommerce_product-variants.md)	 - Product variants commands

