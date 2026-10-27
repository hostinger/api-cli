## hostinger mail logs

Logs commands

```
hostinger mail logs [flags]
```

### Options

```
  -h, --help   help for logs
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
* [hostinger mail logs list-access](hostinger_mail_logs_list-access.md)	 - List access logs
* [hostinger mail logs list-action](hostinger_mail_logs_list-action.md)	 - List action logs
* [hostinger mail logs list-inbound](hostinger_mail_logs_list-inbound.md)	 - List inbound logs
* [hostinger mail logs list-mailbox-action](hostinger_mail_logs_list-mailbox-action.md)	 - List mailbox action logs
* [hostinger mail logs list-outbound](hostinger_mail_logs_list-outbound.md)	 - List outbound logs

