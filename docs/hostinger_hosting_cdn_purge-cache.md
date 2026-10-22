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

### Examples

```
  hostinger hosting cdn purge-cache u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for purge-cache
  -y, --yes    Skip the confirmation prompt
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

