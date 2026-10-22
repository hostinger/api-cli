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

### Examples

```
  hostinger hosting cdn disable u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for disable
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

* [hostinger hosting cdn](hostinger_hosting_cdn.md)	 - CDN commands

