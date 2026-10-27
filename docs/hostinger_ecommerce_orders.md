## hostinger ecommerce orders

Orders commands

```
hostinger ecommerce orders [flags]
```

### Options

```
  -h, --help   help for orders
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
* [hostinger ecommerce orders cancel](hostinger_ecommerce_orders_cancel.md)	 - Cancel an order
* [hostinger ecommerce orders fulfil](hostinger_ecommerce_orders_fulfil.md)	 - Fulfil an order
* [hostinger ecommerce orders list-store](hostinger_ecommerce_orders_list-store.md)	 - List store orders
* [hostinger ecommerce orders retrieve](hostinger_ecommerce_orders_retrieve.md)	 - Retrieve an order

