## hostinger wordpress installations

Installations commands

```
hostinger wordpress installations [flags]
```

### Options

```
  -h, --help   help for installations
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

* [hostinger wordpress](hostinger_wordpress.md)	 - WordPress commands
* [hostinger wordpress installations check-if-are-valid](hostinger_wordpress_installations_check-if-are-valid.md)	 - Check if WordPress installations are valid
* [hostinger wordpress installations delete](hostinger_wordpress_installations_delete.md)	 - Delete WordPress installation
* [hostinger wordpress installations detect](hostinger_wordpress_installations_detect.md)	 - Detect WordPress installations
* [hostinger wordpress installations import-website](hostinger_wordpress_installations_import-website.md)	 - Import WordPress website
* [hostinger wordpress installations install](hostinger_wordpress_installations_install.md)	 - Install WordPress
* [hostinger wordpress installations jwt-token](hostinger_wordpress_installations_jwt-token.md)	 - Get installation JWT token
* [hostinger wordpress installations list](hostinger_wordpress_installations_list.md)	 - List WordPress installations
* [hostinger wordpress installations list-core-updates](hostinger_wordpress_installations_list-core-updates.md)	 - List available WordPress core updates
* [hostinger wordpress installations show-core-version](hostinger_wordpress_installations_show-core-version.md)	 - Show WordPress core version
* [hostinger wordpress installations update-core](hostinger_wordpress_installations_update-core.md)	 - Update WordPress core

