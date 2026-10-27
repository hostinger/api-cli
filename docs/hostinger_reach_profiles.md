## hostinger reach profiles

Profiles commands

```
hostinger reach profiles [flags]
```

### Options

```
  -h, --help   help for profiles
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
* [hostinger reach profiles connect-sending-domain](hostinger_reach_profiles_connect-sending-domain.md)	 - Connect sending domain
* [hostinger reach profiles connected-sending-domain](hostinger_reach_profiles_connected-sending-domain.md)	 - Get connected sending domain
* [hostinger reach profiles delete-sending-domain](hostinger_reach_profiles_delete-sending-domain.md)	 - Delete sending domain
* [hostinger reach profiles domain-dns-status](hostinger_reach_profiles_domain-dns-status.md)	 - Get profile domain DNS status
* [hostinger reach profiles list](hostinger_reach_profiles_list.md)	 - List Profiles
* [hostinger reach profiles list-plan-feature-access](hostinger_reach_profiles_list-plan-feature-access.md)	 - List plan feature access
* [hostinger reach profiles remaining-plan-limits](hostinger_reach_profiles_remaining-plan-limits.md)	 - Get remaining plan limits
* [hostinger reach profiles settings](hostinger_reach_profiles_settings.md)	 - Get profile settings
* [hostinger reach profiles update-settings](hostinger_reach_profiles_update-settings.md)	 - Update profile settings

