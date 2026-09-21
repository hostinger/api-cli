## hostinger agency-hosting ssl install-website

Install website SSL

### Synopsis

Starts a Let's Encrypt certificate setup for the domain and returns at once; the setup runs in
the background. `Get website SSL status` reports `installing` while it runs, then `active` or
`failed`; the `ssl_setup` entry of `List website processes` shows the same progress.

Returns 422 when the domain already has a platform certificate that is not expired, when a
certificate process is recorded for the domain (a failed setup counts until it is cleaned up),
or when the domain hit its limit of three setups per seven days. Returns 429 when the same
domain was requested less than a minute ago, 403 when the website is suspended or locked, and
404 when the website or the domain does not exist.

```
hostinger agency-hosting ssl install-website <website_uid> <domain> [flags]
```

### Options

```
  -h, --help   help for install-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

