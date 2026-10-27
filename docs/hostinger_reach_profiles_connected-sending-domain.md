## hostinger reach profiles connected-sending-domain

Get connected sending domain

### Synopsis

Get the sending domain connected to the profile, its verification status and any suspended
sender addresses.

Campaigns only go out once a domain is connected and active, so this is the cheapest way to
check that precondition before building one. A profile with no domain connected returns the
same shape with every field set to `null`. For the individual MX, SPF, DKIM and DMARC records
behind the status, use the DNS status endpoint.

```
hostinger reach profiles connected-sending-domain <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles connected-sending-domain 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for connected-sending-domain
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

