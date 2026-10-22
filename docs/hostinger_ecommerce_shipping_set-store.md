## hostinger ecommerce shipping set-store

Set store shipping

### Synopsis

Set the flat-rate shipping price for a store, creating the shipping zone if it does not exist yet.

```
hostinger ecommerce shipping set-store <store_id> [flags]
```

### Examples

```
  hostinger ecommerce shipping set-store store_01J8Z5F8W9K8M4A7B3C2D1E0FG --price 500
```

### Options

```
  -h, --help        help for set-store
      --price int   Flat shipping rate in the smallest currency unit (e.g. cents). Use 0 for free shipping.
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

* [hostinger ecommerce shipping](hostinger_ecommerce_shipping.md)	 - Shipping commands

