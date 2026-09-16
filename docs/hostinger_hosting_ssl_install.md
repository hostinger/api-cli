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

```
hostinger hosting ssl install <username> <domain> [flags]
```

### Options

```
  -h, --help   help for install
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

