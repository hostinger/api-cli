## hostinger ecommerce payments list-store-providers

List store payment providers

### Synopsis

List a store's payment providers, split into providers already connected to the store and
gateways available to install. Never exposes gateway credentials, secrets, or configuration.

```
hostinger ecommerce payments list-store-providers <store_id> [flags]
```

### Options

```
  -h, --help                           help for list-store-providers
      --include-currency-unsupported   Include gateways that do not support the store currency in the available list.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce payments](hostinger_ecommerce_payments.md)	 - Payments commands

