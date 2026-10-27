## hostinger agency-hosting ssl reinstall-website

Reinstall website SSL

### Synopsis

Replaces the Let's Encrypt certificate of the domain: the current platform certificate, when
one is recorded, is revoked and removed, then a new setup starts in the background. Returns at
once; `Get website SSL status` reports `installing` while it runs, then `active` or `failed`.

Returns 422 for free subdomains, when a certificate process is recorded for the domain (a
failed setup counts until it is cleaned up), or when the domain hit its limit of three setups
per seven days. Returns 429 when the same domain was requested less than a minute ago, and 403
when the website is suspended or locked, and 404 when the website or the domain does not exist.

```
hostinger agency-hosting ssl reinstall-website <website_uid> <domain> [flags]
```

### Examples

```
  hostinger agency-hosting ssl reinstall-website zpwlGlp19 mydomain.tld
```

### Options

```
  -h, --help   help for reinstall-website
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

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

