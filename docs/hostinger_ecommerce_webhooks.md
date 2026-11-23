## hostinger ecommerce webhooks

Webhooks commands

```
hostinger ecommerce webhooks [flags]
```

### Options

```
  -h, --help[=true|false]   help for webhooks
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce](hostinger_ecommerce.md)	 - Ecommerce commands
* [hostinger ecommerce webhooks create](hostinger_ecommerce_webhooks_create.md)	 - Create an ecommerce webhook
* [hostinger ecommerce webhooks delete](hostinger_ecommerce_webhooks_delete.md)	 - Delete an ecommerce webhook
* [hostinger ecommerce webhooks get](hostinger_ecommerce_webhooks_get.md)	 - Get an ecommerce webhook
* [hostinger ecommerce webhooks list](hostinger_ecommerce_webhooks_list.md)	 - List ecommerce webhooks
* [hostinger ecommerce webhooks rotate-secret](hostinger_ecommerce_webhooks_rotate-secret.md)	 - Rotate an ecommerce webhook secret
* [hostinger ecommerce webhooks update](hostinger_ecommerce_webhooks_update.md)	 - Update an ecommerce webhook

