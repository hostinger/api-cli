## hostinger hosting ssl install

Install SSL

### Synopsis

Requests a lifetime SSL certificate for the website. The installation runs in the background;
`Get SSL status` reports `active` or `failed` when it ends. An `active` lifetime certificate
does not block the request: a new installation is requested, which is how a certificate is
reinstalled.

Returns 422 for free subdomains (their certificate is managed by the platform), while an
installation is `installing` or `waiting_for_retry`, when the website's certificate was
revoked (it cannot be reissued), and when an uploaded custom certificate is installed; that
one has to be uninstalled first.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting ssl install [domain] [flags]
```

### Examples

```
  hostinger hosting ssl install mydomain.tld
```

### Options

```
  -h, --help   help for install
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

