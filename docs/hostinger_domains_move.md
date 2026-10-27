## hostinger domains move

Move commands

```
hostinger domains move [flags]
```

### Options

```
  -h, --help   help for move
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

* [hostinger domains](hostinger_domains.md)	 - Domains commands
* [hostinger domains move accept-incoming](hostinger_domains_move_accept-incoming.md)	 - Accept incoming domain move
* [hostinger domains move cancel-outgoing](hostinger_domains_move_cancel-outgoing.md)	 - Cancel outgoing domain move
* [hostinger domains move incoming](hostinger_domains_move_incoming.md)	 - Get incoming domain move
* [hostinger domains move incoming-list](hostinger_domains_move_incoming-list.md)	 - Get incoming domain move list
* [hostinger domains move outgoing](hostinger_domains_move_outgoing.md)	 - Get outgoing domain move
* [hostinger domains move outgoing-list](hostinger_domains_move_outgoing-list.md)	 - Get outgoing domain move list
* [hostinger domains move reject-incoming](hostinger_domains_move_reject-incoming.md)	 - Reject incoming domain move
* [hostinger domains move start-outgoing](hostinger_domains_move_start-outgoing.md)	 - Start outgoing domain move

