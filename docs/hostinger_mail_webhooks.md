## hostinger mail webhooks

Webhooks commands

```
hostinger mail webhooks [flags]
```

### Options

```
  -h, --help   help for webhooks
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

* [hostinger mail](hostinger_mail.md)	 - Mail commands
* [hostinger mail webhooks create](hostinger_mail_webhooks_create.md)	 - Create webhook
* [hostinger mail webhooks delete](hostinger_mail_webhooks_delete.md)	 - Delete webhook
* [hostinger mail webhooks get](hostinger_mail_webhooks_get.md)	 - Get webhook
* [hostinger mail webhooks list](hostinger_mail_webhooks_list.md)	 - List webhooks
* [hostinger mail webhooks list-delivery-logs](hostinger_mail_webhooks_list-delivery-logs.md)	 - List webhook delivery logs
* [hostinger mail webhooks regenerate-secret](hostinger_mail_webhooks_regenerate-secret.md)	 - Regenerate webhook secret
* [hostinger mail webhooks test](hostinger_mail_webhooks_test.md)	 - Test webhook
* [hostinger mail webhooks update](hostinger_mail_webhooks_update.md)	 - Update webhook

