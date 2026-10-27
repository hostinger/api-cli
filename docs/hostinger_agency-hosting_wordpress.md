## hostinger agency-hosting wordpress

WordPress commands

```
hostinger agency-hosting wordpress [flags]
```

### Options

```
  -h, --help   help for wordpress
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

* [hostinger agency-hosting](hostinger_agency-hosting.md)	 - Agency Hosting commands
* [hostinger agency-hosting wordpress change-version](hostinger_agency-hosting_wordpress_change-version.md)	 - Change WordPress version
* [hostinger agency-hosting wordpress list-versions](hostinger_agency-hosting_wordpress_list-versions.md)	 - List available WordPress versions
* [hostinger agency-hosting wordpress settings](hostinger_agency-hosting_wordpress_settings.md)	 - Get WordPress settings

