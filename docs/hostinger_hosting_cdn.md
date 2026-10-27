## hostinger hosting cdn

CDN commands

```
hostinger hosting cdn [flags]
```

### Options

```
  -h, --help   help for cdn
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

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting cdn disable](hostinger_hosting_cdn_disable.md)	 - Disable CDN
* [hostinger hosting cdn enable](hostinger_hosting_cdn_enable.md)	 - Enable CDN
* [hostinger hosting cdn purge-cache](hostinger_hosting_cdn_purge-cache.md)	 - Purge CDN cache
* [hostinger hosting cdn status](hostinger_hosting_cdn_status.md)	 - Get CDN status

