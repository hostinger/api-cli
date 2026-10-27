## hostinger agency-hosting php update-website-version

Update website PHP version

### Synopsis

Switches an Agency Plan website to a different PHP version. Call the available versions endpoint first to see which versions can be selected. The website restarts on the new version, so requests served during the switch may fail and code that is incompatible with the target version will break.

```
hostinger agency-hosting php update-website-version <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php update-website-version zpwlGlp19 --version 8.2
```

### Options

```
  -h, --help             help for update-website-version
      --version string   PHP version to switch the website to, as major.minor. Must be one of the versions returned by the available versions endpoint.
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

