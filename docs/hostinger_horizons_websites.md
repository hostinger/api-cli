## hostinger horizons websites

Websites commands

```
hostinger horizons websites [flags]
```

### Options

```
  -h, --help   help for websites
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

* [hostinger horizons](hostinger_horizons.md)	 - Horizons commands
* [hostinger horizons websites clone](hostinger_horizons_websites_clone.md)	 - Clone website
* [hostinger horizons websites create](hostinger_horizons_websites_create.md)	 - Create website
* [hostinger horizons websites edit](hostinger_horizons_websites_edit.md)	 - Edit website
* [hostinger horizons websites get](hostinger_horizons_websites_get.md)	 - Get website
* [hostinger horizons websites list](hostinger_horizons_websites_list.md)	 - Get website list
* [hostinger horizons websites publish](hostinger_horizons_websites_publish.md)	 - Publish website

