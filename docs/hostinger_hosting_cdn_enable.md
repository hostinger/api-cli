## hostinger hosting cdn enable

Enable CDN

### Synopsis

Turns on the Hostinger CDN for the website domain, which speeds up page loads and reduces the
load on the server. Works for the main domain, parked domains and subdomains of the account.

Enabling changes the domain's DNS records and usually takes a few seconds. After that, DNS
propagation can take a few more minutes before visitors are served through the CDN.

Repeating the call is safe, for example after a timeout. When CDN is already enabled, it
applies the DNS change again and refreshes the CDN SSL certificate.

Returns 422 when the account is disabled, while the account is being moved to another
server (retry after the move finishes), or when the CDN provider rejects the change.
Returns 404 when the domain is not a website on the account.

```
hostinger hosting cdn enable <username> <domain> [flags]
```

### Examples

```
  hostinger hosting cdn enable u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for enable
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

