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

### Options

```
  -h, --help   help for reinstall-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

