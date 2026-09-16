## hostinger hosting ssl uninstall

Uninstall SSL

### Synopsis

Removes the SSL certificate assigned to the website, turns the HTTPS redirect off and cancels
a pending installation retry. The website serves plain HTTP until a new installation
completes. `Get SSL status` reports `not_installed` as soon as the call returns; the call also
succeeds when no certificate is assigned, so repeating it is safe.

Returns 422 for free subdomains (their certificate is managed by the platform) and while an
installation is `installing`.

```
hostinger hosting ssl uninstall <username> <domain> [flags]
```

### Options

```
  -h, --help   help for uninstall
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

