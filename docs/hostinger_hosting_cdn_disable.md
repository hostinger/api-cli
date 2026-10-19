## hostinger hosting cdn disable

Disable CDN

### Synopsis

Turns off the Hostinger CDN for the website domain. Once DNS propagation completes, the
website content is served directly from the origin server. Repeating the call is safe; when
CDN is already disabled, it applies the DNS change again.

Returns 422 for platform-managed domains (for example `*.hostingersite.com` and
`*.app-preview.com`), where CDN cannot be disabled, even when CDN is already off. Also returns
422 while the account is being moved to another server (retry after the move finishes), or
when the CDN provider rejects the change. Returns 404 when the domain is not a website on the
account.

```
hostinger hosting cdn disable <username> <domain> [flags]
```

### Options

```
  -h, --help   help for disable
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting cdn](hostinger_hosting_cdn.md)	 - CDN commands

