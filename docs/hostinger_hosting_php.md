## hostinger hosting php

PHP commands

```
hostinger hosting php [flags]
```

### Options

```
  -h, --help   help for php
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
* [hostinger hosting php get](hostinger_hosting_php_get.md)	 - Get PHP details
* [hostinger hosting php info](hostinger_hosting_php_info.md)	 - Get PHP info
* [hostinger hosting php reset-extensions](hostinger_hosting_php_reset-extensions.md)	 - Reset PHP extensions
* [hostinger hosting php update-extensions](hostinger_hosting_php_update-extensions.md)	 - Update PHP extensions
* [hostinger hosting php update-options](hostinger_hosting_php_update-options.md)	 - Update PHP options
* [hostinger hosting php update-version](hostinger_hosting_php_update-version.md)	 - Update PHP version

