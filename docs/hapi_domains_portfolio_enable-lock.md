## hapi domains portfolio enable-lock

Enable domain lock

### Synopsis

Enable domain lock for the domain.

When domain lock is enabled,
the domain cannot be transferred to another registrar without first disabling the lock.

Use this endpoint to secure domains against unauthorized transfers.

```
hapi domains portfolio enable-lock <domain> [flags]
```

### Options

```
  -h, --help   help for enable-lock
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi domains portfolio](hapi_domains_portfolio.md)	 - Portfolio commands

