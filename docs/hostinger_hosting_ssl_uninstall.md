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

### Examples

```
  hostinger hosting ssl uninstall u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for uninstall
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

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

