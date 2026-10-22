## hostinger ecommerce orders fulfil

Fulfil an order

### Synopsis

Create a fulfilment for the order and attach tracking in one call. Omit items to fulfil
every remaining unfulfilled item. Returns the updated order summary.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --items
    [
      {
        "line_item_id": "item_01J8Z5F8W9K8M4A7B3C2D1E0FG",
        "quantity": 2
      }
    ]

```
hostinger ecommerce orders fulfil <store_id> <order_id> [flags]
```

### Examples

```
  hostinger ecommerce orders fulfil store_01J8Z5F8W9K8M4A7B3C2D1E0FG order_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help                     help for fulfil
      --items string             Line items to fulfil. Omit to fulfil every remaining unfulfilled item. (JSON, @file or -; shape in the help above)
      --notify-customer          Whether to email the customer about the fulfilment. Defaults to true.
      --tracking-number string   Carrier tracking number for the shipment.
      --tracking-url string      Public tracking URL for the shipment. Requires tracking_number.
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

* [hostinger ecommerce orders](hostinger_ecommerce_orders.md)	 - Orders commands

