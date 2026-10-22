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

### Examples

```
  hostinger hosting cache toggle-cacheless u123456789 mydomain.tld --enabled=true
```

### Options

```
      --enabled   Turn development (cacheless) mode on (true) or off (false) for the website.
  -h, --help      help for toggle-cacheless
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

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

