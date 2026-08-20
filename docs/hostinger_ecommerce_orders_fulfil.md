## hostinger ecommerce orders fulfil

Fulfil an order

### Synopsis

Create a fulfilment for the order and attach tracking in one call. Omit items to fulfil
every remaining unfulfilled item. Returns the updated order summary.

```
hostinger ecommerce orders fulfil <store_id> <order_id> [flags]
```

### Options

```
  -h, --help                     help for fulfil
      --items string             Line items to fulfil. Omit to fulfil every remaining unfulfilled item. (JSON)
      --notify-customer          Whether to email the customer about the fulfilment. Defaults to true.
      --tracking-number string   Carrier tracking number for the shipment.
      --tracking-url string      Public tracking URL for the shipment. Requires tracking_number.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce orders](hostinger_ecommerce_orders.md)	 - Orders commands

