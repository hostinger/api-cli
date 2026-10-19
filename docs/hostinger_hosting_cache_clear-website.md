## hostinger hosting cache clear-website

Clear website cache

### Synopsis

Permanently clears all server-side cache for the website at once. Use it when content was
updated and needs to be visible immediately, or after making major changes.

Also purges the Hostinger CDN cache when CDN is enabled on the website. To purge only the CDN
cache and leave the server cache untouched, use `Purge CDN cache`. For a WordPress
installation living in a subdirectory, pass the directory query parameter to clear its cache.

```
hostinger hosting cache clear-website <username> <domain> [flags]
```

### Options

```
      --directory string   Directory of the website installation to clear, relative to the website root.
                           Defaults to the website root.
  -h, --help               help for clear-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

