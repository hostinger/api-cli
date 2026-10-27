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

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting cdn enable [domain] [flags]
```

### Examples

```
  hostinger hosting cdn enable mydomain.tld
```

### Options

```
  -h, --help   help for enable
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting cdn](hostinger_hosting_cdn.md)	 - CDN commands

