## hostinger vps post-install-scripts

Post-install scripts commands

```
hostinger vps post-install-scripts [flags]
```

### Options

```
  -h, --help   help for post-install-scripts
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
* [hostinger vps post-install-scripts create](hostinger_vps_post-install-scripts_create.md)	 - Create post-install script
* [hostinger vps post-install-scripts delete](hostinger_vps_post-install-scripts_delete.md)	 - Delete post-install script
* [hostinger vps post-install-scripts get](hostinger_vps_post-install-scripts_get.md)	 - Get post-install script
* [hostinger vps post-install-scripts list](hostinger_vps_post-install-scripts_list.md)	 - Get post-install scripts
* [hostinger vps post-install-scripts update](hostinger_vps_post-install-scripts_update.md)	 - Update post-install script

