## hostinger ecommerce orders cancel

Cancel an order

### Synopsis

Cancel the order and optionally email the customer. Returns the updated order summary.

```
hostinger ecommerce orders cancel <store_id> <order_id> [flags]
```

### Options

```
  -h, --help              help for cancel
      --notify-customer   Whether to email the customer about the cancellation. Defaults to true.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce orders](hostinger_ecommerce_orders.md)	 - Orders commands

