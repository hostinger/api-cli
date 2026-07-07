## hostinger hosting cache disable-website

Disable website cache

### Synopsis

Turns off server-side caching for the website until it is enabled again. May impact performance.
Use it when experiencing cache-related issues; to temporarily bypass caching while developing
or debugging, prefer enabling cacheless mode instead.

Does nothing if caching is already disabled.

```
hostinger hosting cache disable-website <username> <domain> [flags]
```

### Options

```
  -h, --help   help for disable-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

