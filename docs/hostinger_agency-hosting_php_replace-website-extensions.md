## hostinger agency-hosting php replace-website-extensions

Replace website PHP extensions

### Synopsis

Replaces the set of PHP extensions enabled on an Agency Plan website with the ones provided. Any toggleable extension not in the request is disabled, so call the extensions endpoint first and send the full desired set. Extensions compiled into PHP, reported with the "built-in" state, are always active and are unaffected.

```
hostinger agency-hosting php replace-website-extensions <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php replace-website-extensions zpwlGlp19 --extensions bz2,zip
```

### Options

```
      --extensions strings   Extension names, exactly as returned by the extensions endpoint.
  -h, --help                 help for replace-website-extensions
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

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

