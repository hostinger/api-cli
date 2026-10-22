## hostinger hosting cache toggle-website

Toggle website cache

### Synopsis

Turns server-side caching for the website on or off, based on the enabled flag. Enable it for
faster page loads, reduced server load, and improved user experience; recommended for production
websites. Disabling may impact performance; to temporarily bypass caching while developing or
debugging, prefer toggling cacheless mode instead.

Does nothing if caching is already in the requested state.

```
hostinger hosting cache toggle-website <username> <domain> [flags]
```

### Examples

```
  hostinger hosting cache toggle-website u123456789 mydomain.tld --enabled=true
```

### Options

```
      --enabled   Turn server-side caching on (true) or off (false) for the website.
  -h, --help      help for toggle-website
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

