## hostinger domains move start-outgoing

Start outgoing domain move

### Synopsis

Initiate a move of a specified domain to another Hostinger account.

The receiving account has to already exist and accept the move before the domain changes hands.

The domain must be active. The subscription it belongs to is resolved automatically,
and the request is rejected with a 404 status code when the domain has no domain
subscription of its own.

Domains protected by premium protection require an additional verification step,
such requests are rejected with a 428 status code.

Use this endpoint to hand a domain over to another Hostinger user.

The domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger domains move start-outgoing [domain] [flags]
```

### Examples

```
  hostinger domains move start-outgoing mydomain.tld --new-customer-email new-owner@example.com
```

### Options

```
  -h, --help                        help for start-outgoing
      --new-customer-email string   Email address of the Hostinger account receiving the domain
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

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

