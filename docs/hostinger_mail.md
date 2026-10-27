## hostinger mail

Mail commands

```
hostinger mail [flags]
```

### Options

```
  -h, --help   help for mail
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

* [hostinger](hostinger.md)	 - Hostinger API Command Line Interface
* [hostinger mail aliases](hostinger_mail_aliases.md)	 - Aliases commands
* [hostinger mail api-tokens](hostinger_mail_api-tokens.md)	 - API Tokens commands
* [hostinger mail autoreplies](hostinger_mail_autoreplies.md)	 - Autoreplies commands
* [hostinger mail catchalls](hostinger_mail_catchalls.md)	 - Catchalls commands
* [hostinger mail forwarders](hostinger_mail_forwarders.md)	 - Forwarders commands
* [hostinger mail logs](hostinger_mail_logs.md)	 - Logs commands
* [hostinger mail mailboxes](hostinger_mail_mailboxes.md)	 - Mailboxes commands
* [hostinger mail orders](hostinger_mail_orders.md)	 - Orders commands
* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

