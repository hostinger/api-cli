## hostinger agency-hosting ssl uninstall-website

Uninstall website SSL

### Synopsis

Removes the platform-issued Let's Encrypt certificate of the domain: the certificate is revoked
and deleted before the response, so the domain is no longer served with a platform certificate
until a new setup completes. Also succeeds when the domain has no platform certificate to
remove. Uploaded (custom) certificates are not affected.

Returns 422 when a certificate process is recorded for the domain (a failed setup counts until
it is cleaned up), 429 when the same domain was requested less than a minute ago, and 403 when
the website is suspended or locked, and 404 when the website or the domain does not exist.

```
hostinger agency-hosting ssl uninstall-website <website_uid> <domain> [flags]
```

### Examples

```
  hostinger agency-hosting ssl uninstall-website zpwlGlp19 mydomain.tld
```

### Options

```
  -h, --help   help for uninstall-website
  -y, --yes    Skip the confirmation prompt
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

