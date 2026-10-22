## hostinger ecommerce payments create-provider-connect-link

Create a payment provider connect link

### Synopsis

Create an onboarding link for connecting a payment gateway to the store. Returns the gateway
onboarding URL for the merchant to open and a deep-link into the store admin.

```
hostinger ecommerce payments create-provider-connect-link <store_id> <provider_id> [flags]
```

### Examples

```
  hostinger ecommerce payments create-provider-connect-link store_01J8Z5F8W9K8M4A7B3C2D1E0FG stripe
```

### Options

```
  -h, --help   help for create-provider-connect-link
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

