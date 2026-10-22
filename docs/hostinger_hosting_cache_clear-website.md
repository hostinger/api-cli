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

### Examples

```
  hostinger hosting cache clear-website u123456789 mydomain.tld
```

### Options

```
      --directory string   Directory of the website installation to clear, relative to the website root. Defaults to the website root.
  -h, --help               help for clear-website
  -y, --yes                Skip the confirmation prompt
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

