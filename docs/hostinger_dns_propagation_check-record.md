## hostinger dns propagation check-record

Check DNS record propagation

### Synopsis

Check whether a DNS record of a domain or host name has propagated to public resolvers around the world.

`authoritative` shows what the domain's authoritative nameservers serve now. It is the best signal
of the current record value. Its TTL tells how long resolvers may keep serving the old value.

`vantage_points` shows what public resolvers at different locations return. `is_propagated` is true
when the location returns the same records as the authoritative nameservers. For GeoDNS domains,
whose nameservers answer differently by region, true means the location returned a record.
A records served from the Hostinger CDN count as propagated. When every responding authoritative
nameserver returns no records, `is_propagated` is true.

Works for any domain or host name, either a domain or a subdomain. Read-only: it changes nothing.

A 404 response with error code 2025 can also mean the check could not be completed, for example
when the DNS checker is busy or timed out. It is safe to retry after a short delay.

The domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger dns propagation check-record [domain] <type> [flags]
```

### Examples

```
  hostinger dns propagation check-record www.mydomain.tld A
```

### Options

```
  -h, --help[=true|false]   help for check-record
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger dns propagation](hostinger_dns_propagation.md)	 - Propagation commands

