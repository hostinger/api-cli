## hostinger agency-hosting php

PHP commands

```
hostinger agency-hosting php [flags]
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

* [hostinger agency-hosting](hostinger_agency-hosting.md)	 - Agency Hosting commands
* [hostinger agency-hosting php list-extensions-for-website](hostinger_agency-hosting_php_list-extensions-for-website.md)	 - List PHP extensions for a website
* [hostinger agency-hosting php list-options-for-website](hostinger_agency-hosting_php_list-options-for-website.md)	 - List PHP options for a website
* [hostinger agency-hosting php list-versions-for-order](hostinger_agency-hosting_php_list-versions-for-order.md)	 - List available PHP versions for an order
* [hostinger agency-hosting php list-versions-for-website](hostinger_agency-hosting_php_list-versions-for-website.md)	 - List available PHP versions for a website
* [hostinger agency-hosting php replace-website-extensions](hostinger_agency-hosting_php_replace-website-extensions.md)	 - Replace website PHP extensions
* [hostinger agency-hosting php replace-website-options](hostinger_agency-hosting_php_replace-website-options.md)	 - Replace website PHP options
* [hostinger agency-hosting php update-website-version](hostinger_agency-hosting_php_update-website-version.md)	 - Update website PHP version

