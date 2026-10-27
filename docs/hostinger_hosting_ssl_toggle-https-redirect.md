## hostinger hosting ssl toggle-https-redirect

Toggle HTTPS redirect

### Synopsis

Turns the HTTP to HTTPS redirect of the website on or off, based on `is_enabled`. Does
nothing when the redirect is already in the requested state. Turning it on requires an
installed certificate (`status` `active` or `expired` on `Get SSL status`) and returns 422
when there is none; turning it off is always accepted.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting ssl toggle-https-redirect [domain] [flags]
```

### Examples

```
  hostinger hosting ssl toggle-https-redirect mydomain.tld --is-enabled=true
```

### Options

```
  -h, --help         help for toggle-https-redirect
      --is-enabled   Turn the HTTP to HTTPS redirect on (true) or off (false) for the website.
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

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

