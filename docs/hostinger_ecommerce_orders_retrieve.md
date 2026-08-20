## hostinger ecommerce orders retrieve

Retrieve an order

### Synopsis

Retrieve one order in full: line items (each with the id the fulfil endpoint needs),
addresses, the totals breakdown and fulfilments with tracking. Amounts are in the
smallest currency unit.

```
hostinger ecommerce orders retrieve <store_id> <order_id> [flags]
```

### Options

```
  -h, --help   help for retrieve
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce orders](hostinger_ecommerce_orders.md)	 - Orders commands

