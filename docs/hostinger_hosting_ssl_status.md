## hostinger hosting ssl status

Get SSL status

### Synopsis

Returns the SSL state of the website: the certificate `status` and `provider`, whether the
certificate is a lifetime one managed by the platform, whether HTTP requests are redirected to
HTTPS, when the certificate stops being valid and the last installation error.

`installing` and `waiting_for_retry` mean an installation is in progress. `failed` means the
last installation gave up, or the website was not updated for 60 minutes while `installing`;
`last_error` holds the reason when it is a known message, otherwise it is null. `expired`
means the assigned certificate's validity has ended. `not_installed` means no certificate is
assigned. Free subdomains use a platform-managed certificate: with no installation recorded
they report `active` with `provider` and `expires_at` null.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting ssl status [domain] [flags]
```

### Examples

```
  hostinger hosting ssl status mydomain.tld
```

### Options

```
  -h, --help   help for status
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

