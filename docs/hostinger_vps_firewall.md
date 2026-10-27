## hostinger vps firewall

Firewall commands

```
hostinger vps firewall [flags]
```

### Options

```
  -h, --help   help for firewall
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

* [hostinger vps](hostinger_vps.md)	 - VPS commands
* [hostinger vps firewall activate](hostinger_vps_firewall_activate.md)	 - Activate firewall
* [hostinger vps firewall create](hostinger_vps_firewall_create.md)	 - Create new firewall
* [hostinger vps firewall create-rule](hostinger_vps_firewall_create-rule.md)	 - Create firewall rule
* [hostinger vps firewall deactivate](hostinger_vps_firewall_deactivate.md)	 - Deactivate firewall
* [hostinger vps firewall delete](hostinger_vps_firewall_delete.md)	 - Delete firewall
* [hostinger vps firewall delete-rule](hostinger_vps_firewall_delete-rule.md)	 - Delete firewall rule
* [hostinger vps firewall get](hostinger_vps_firewall_get.md)	 - Get firewall details
* [hostinger vps firewall list](hostinger_vps_firewall_list.md)	 - Get firewall list
* [hostinger vps firewall replace-all-rules-in-group](hostinger_vps_firewall_replace-all-rules-in-group.md)	 - Replace all firewall rules in group
* [hostinger vps firewall sync](hostinger_vps_firewall_sync.md)	 - Sync firewall
* [hostinger vps firewall sync-to-all-assigned-v-ms](hostinger_vps_firewall_sync-to-all-assigned-v-ms.md)	 - Sync firewall to all assigned VMs
* [hostinger vps firewall update-rule](hostinger_vps_firewall_update-rule.md)	 - Update firewall rule

