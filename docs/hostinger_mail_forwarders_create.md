## hostinger mail forwarders create

Create forwarder

### Synopsis

Create a forwarder from the given mailbox to the destination address.
The destination receives a confirmation email and forwarding becomes
active only after it is confirmed.

```
hostinger mail forwarders create <mailbox-id> [flags]
```

### Examples

```
  hostinger mail forwarders create AC1a2b3c4d5e6f7g --destination jane@example.org
```

### Options

```
      --destination string     Email address the messages will be forwarded to
  -h, --help                   help for create
      --is-keep-copy-enabled   Whether to keep a copy of forwarded messages in the mailbox. Defaults to false.
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

* [hostinger mail forwarders](hostinger_mail_forwarders.md)	 - Forwarders commands

