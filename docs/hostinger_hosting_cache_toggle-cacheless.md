## hostinger hosting cache toggle-cacheless

Toggle cacheless mode

### Synopsis

Turns development (cacheless) mode on or off, based on the enabled flag. When enabled, nothing
is cached, effectively turning off all caching for the website; use it while actively developing,
testing changes, debugging issues, or when real-time updates must be visible. Disable it after
finishing development work to restore the performance benefits of caching.

```
hostinger hosting cache toggle-cacheless <username> <domain> [flags]
```

### Options

```
      --enabled   Turn development (cacheless) mode on (true) or off (false) for the website.
  -h, --help      help for toggle-cacheless
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

