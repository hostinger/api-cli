## hostinger ecommerce payments enable-manual-method

Enable manual payment method

### Synopsis

Enable a manual payment method so the store can accept orders without an online payment provider.

```
hostinger ecommerce payments enable-manual-method <store_id> [flags]
```

### Examples

```
  hostinger ecommerce payments enable-manual-method store_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help           help for enable-manual-method
      --title string   Optional display name shown to customers at checkout.
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce payments](hostinger_ecommerce_payments.md)	 - Payments commands

