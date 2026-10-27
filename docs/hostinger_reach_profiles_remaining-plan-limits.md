## hostinger reach profiles remaining-plan-limits

Get remaining plan limits

### Synopsis

Get how much of the plan is left for the current period.

Two things to keep in mind before you build alerting on this. The period is a calendar month
rather than a billing anniversary, so the counters reset on the 1st no matter when the
subscription started. And usage is tracked per order, so every profile on the same order shares
one pool and reports the same numbers here. Only the current period is available, past usage is
not kept.

```
hostinger reach profiles remaining-plan-limits <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles remaining-plan-limits 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for remaining-plan-limits
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

