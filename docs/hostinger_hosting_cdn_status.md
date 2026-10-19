## hostinger hosting cdn status

Get CDN status

### Synopsis

Returns whether the Hostinger CDN is enabled for the website domain. Works for the main
domain, parked domains and subdomains of the account.

The domain must be the exact website name as listed on the account, without a `www.` prefix.
Returns 404 when the domain is not a website on the account.

```
hostinger hosting cdn status <username> <domain> [flags]
```

### Options

```
  -h, --help   help for status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cdn](hostinger_hosting_cdn.md)	 - CDN commands

