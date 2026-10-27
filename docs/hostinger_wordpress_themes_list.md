## hostinger wordpress themes list

List WordPress themes

### Synopsis

List WordPress themes available to install.

Use the returned `slug` values with
POST /api/hosting/v1/accounts/{username}/wordpress/{software}/themes/install.

```
hostinger wordpress themes list [flags]
```

### Examples

```
  hostinger wordpress themes list
```

### Options

```
  -h, --help            help for list
      --order-id int    Optionally scope themes to a specific order.
      --search string   Search term to match against theme names.
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

