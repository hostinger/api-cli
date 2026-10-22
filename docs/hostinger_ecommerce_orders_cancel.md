## hostinger ecommerce orders cancel

Cancel an order

### Synopsis

Cancel the order and optionally email the customer. Returns the updated order summary.

```
hostinger ecommerce orders cancel <store_id> <order_id> [flags]
```

### Examples

```
  hostinger ecommerce orders cancel store_01J8Z5F8W9K8M4A7B3C2D1E0FG order_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help              help for cancel
      --notify-customer   Whether to email the customer about the cancellation. Defaults to true.
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

