## hostinger hosting cdn purge-cache

Purge CDN cache

### Synopsis

Purges the website's content from the Hostinger CDN edge cache only. The origin server cache
is not touched; to clear both, use `Clear website cache`.

Returns 422 when CDN is not enabled on the website. Returns 404 when the domain is not a
website on the account.

```
hostinger hosting cdn purge-cache <username> <domain> [flags]
```

### Options

```
  -h, --help   help for purge-cache
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cdn](hostinger_hosting_cdn.md)	 - CDN commands

