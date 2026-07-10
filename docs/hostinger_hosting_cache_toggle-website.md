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

### Options

```
      --enabled   Turn server-side caching on (true) or off (false) for the website.
  -h, --help      help for toggle-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

