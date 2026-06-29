## hostinger billing catalog list

Get catalog item list

### Synopsis

Retrieve catalog items available for order.

Prices in catalog items is displayed as cents (without floating point),
e.g: float `17.99` is displayed as integer `1799`.

Use this endpoint to view available services and pricing before placing orders.

```
hostinger billing catalog list [flags]
```

### Options

```
      --category string   Filter catalog items by category (one of: DOMAIN, VPS)
  -h, --help              help for list
      --name *            Filter catalog items by name. Use * for wildcard search, e.g. `.COM*` to find .com domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger billing catalog](hostinger_billing_catalog.md)	 - Catalog commands

