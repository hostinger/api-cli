## hostinger hosting ssl uninstall

Uninstall SSL

### Synopsis

Removes the SSL certificate assigned to the website, turns the HTTPS redirect off and cancels
a pending installation retry. The website serves plain HTTP until a new installation
completes. `Get SSL status` reports `not_installed` as soon as the call returns; the call also
succeeds when no certificate is assigned, so repeating it is safe.

Returns 422 for free subdomains (their certificate is managed by the platform) and while an
installation is `installing`.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting ssl uninstall [domain] [flags]
```

### Examples

```
  hostinger hosting ssl uninstall mydomain.tld
```

### Options

```
  -h, --help   help for uninstall
  -y, --yes    Skip the confirmation prompt
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

