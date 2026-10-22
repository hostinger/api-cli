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

### Examples

```
  hostinger hosting cdn status u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for status
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

