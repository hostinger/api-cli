## hostinger reach campaigns send

Send campaign

### Synopsis

Send a campaign now to the audience stored on it.

This is irreversible: once accepted, the emails go out and cannot be recalled. Confirm with the
customer before calling. The campaign needs a subject, sender, template and audience first - set
them with the update endpoint. Delivery is asynchronous, so success means accepted, not delivered:
the campaign moves to `sending` and later to `publish`, or to `failed`. Read the campaign to follow
it. The request is rejected when the campaign has no recipients, its sender domain is not verified
or is blocked, the plan's sending limits would be exceeded, or the campaign is already being sent.

```
hostinger reach campaigns send <profile-uuid> <campaign-uuid> [flags]
```

### Examples

```
  hostinger reach campaigns send 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help[=true|false]   help for send
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

