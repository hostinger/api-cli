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

### Options

```
  -h, --help   help for uninstall-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

