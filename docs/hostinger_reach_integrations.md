## hostinger reach integrations

Integrations commands

```
hostinger reach integrations [flags]
```

### Options

```
  -h, --help   help for integrations
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
* [hostinger reach integrations connect-ai-website-builder-site](hostinger_reach_integrations_connect-ai-website-builder-site.md)	 - Connect AI Website Builder site
* [hostinger reach integrations disconnect-ai-website-builder-site](hostinger_reach_integrations_disconnect-ai-website-builder-site.md)	 - Disconnect AI Website Builder site
* [hostinger reach integrations disconnect-ecommerce-store](hostinger_reach_integrations_disconnect-ecommerce-store.md)	 - Disconnect Ecommerce store
* [hostinger reach integrations disconnect-wordpress-site](hostinger_reach_integrations_disconnect-wordpress-site.md)	 - Disconnect WordPress site
* [hostinger reach integrations list-connected-websites](hostinger_reach_integrations_list-connected-websites.md)	 - List connected websites
* [hostinger reach integrations list-ecommerce-stores](hostinger_reach_integrations_list-ecommerce-stores.md)	 - List Ecommerce stores

