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

### Examples

```
  hostinger billing catalog list
```

### Options

```
      --category string   Filter catalog items by category (one of: DOMAIN, VPS, EMAIL, HOSTING)
  -h, --help              help for list
      --name *            Filter catalog items by name. Use * for wildcard search, e.g. `.COM*` to find .com domain
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

* [hostinger billing catalog](hostinger_billing_catalog.md)	 - Catalog commands

