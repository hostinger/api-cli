## hostinger reach integrations list-connected-websites

List connected websites

### Synopsis

Get the websites connected to the profile, each with its `domain` and `platform_type`.

A domain is connected to at most one profile, so to find which plan a site is on, call this
for each profile. Only connected sites are listed, not every website the customer owns, and
eCommerce stores are not included - use the list eCommerce stores endpoint for those. Pass
`domain` to the disconnect WordPress site or disconnect AI Website Builder site endpoint.

```
hostinger reach integrations list-connected-websites <profile-uuid> [flags]
```

### Examples

```
  hostinger reach integrations list-connected-websites 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-connected-websites
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

* [hostinger reach integrations](hostinger_reach_integrations.md)	 - Integrations commands

