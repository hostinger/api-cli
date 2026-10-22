## hostinger domains portfolio disable-lock

Disable domain lock

### Synopsis

Disable domain lock for the domain.

Domain lock needs to be disabled before transferring the domain to another registrar.

Use this endpoint to prepare domains for transfer to other registrars.

```
hostinger domains portfolio disable-lock <domain> [flags]
```

### Examples

```
  hostinger domains portfolio disable-lock mydomain.tld
```

### Options

```
  -h, --help   help for disable-lock
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

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

