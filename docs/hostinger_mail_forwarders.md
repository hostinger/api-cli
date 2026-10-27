## hostinger mail forwarders

Forwarders commands

```
hostinger mail forwarders [flags]
```

### Options

```
  -h, --help   help for forwarders
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
* [hostinger mail forwarders create](hostinger_mail_forwarders_create.md)	 - Create forwarder
* [hostinger mail forwarders delete](hostinger_mail_forwarders_delete.md)	 - Delete forwarder
* [hostinger mail forwarders list](hostinger_mail_forwarders_list.md)	 - List forwarders
* [hostinger mail forwarders resend-confirmation](hostinger_mail_forwarders_resend-confirmation.md)	 - Resend forwarder confirmation
* [hostinger mail forwarders update-keep-copy-setting](hostinger_mail_forwarders_update-keep-copy-setting.md)	 - Update forwarder keep-copy setting

