## hostinger reach campaigns

Campaigns commands

```
hostinger reach campaigns [flags]
```

### Options

```
  -h, --help   help for campaigns
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

* [hostinger reach](hostinger_reach.md)	 - Reach commands
* [hostinger reach campaigns create-draft](hostinger_reach_campaigns_create-draft.md)	 - Create a draft campaign
* [hostinger reach campaigns get](hostinger_reach_campaigns_get.md)	 - Get campaign details
* [hostinger reach campaigns list](hostinger_reach_campaigns_list.md)	 - List campaigns
* [hostinger reach campaigns performance](hostinger_reach_campaigns_performance.md)	 - Get campaign performance

