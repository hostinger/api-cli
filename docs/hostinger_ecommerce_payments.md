## hostinger ecommerce payments

Payments commands

```
hostinger ecommerce payments [flags]
```

### Options

```
  -h, --help   help for payments
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

* [hostinger ecommerce](hostinger_ecommerce.md)	 - Ecommerce commands
* [hostinger ecommerce payments create-provider-connect-link](hostinger_ecommerce_payments_create-provider-connect-link.md)	 - Create a payment provider connect link
* [hostinger ecommerce payments enable-manual-method](hostinger_ecommerce_payments_enable-manual-method.md)	 - Enable manual payment method
* [hostinger ecommerce payments list-store-providers](hostinger_ecommerce_payments_list-store-providers.md)	 - List store payment providers

