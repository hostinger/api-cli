## hostinger ecommerce payments list-store-providers

List store payment providers

### Synopsis

List a store's payment providers, split into providers already connected to the store and
gateways available to install. Never exposes gateway credentials, secrets, or configuration.

```
hostinger ecommerce payments list-store-providers <store_id> [flags]
```

### Examples

```
  hostinger ecommerce payments list-store-providers store_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help                           help for list-store-providers
      --include-currency-unsupported   Include gateways that do not support the store currency in the available list.
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

* [hostinger ecommerce payments](hostinger_ecommerce_payments.md)	 - Payments commands

