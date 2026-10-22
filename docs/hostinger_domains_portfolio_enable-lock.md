## hostinger domains portfolio enable-lock

Enable domain lock

### Synopsis

Enable domain lock for the domain.

When domain lock is enabled,
the domain cannot be transferred to another registrar without first disabling the lock.

Use this endpoint to secure domains against unauthorized transfers.

```
hostinger domains portfolio enable-lock <domain> [flags]
```

### Examples

```
  hostinger domains portfolio enable-lock mydomain.tld
```

### Options

```
  -h, --help   help for enable-lock
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

