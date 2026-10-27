## hostinger hosting cache status

Get cache status

### Synopsis

Returns the caching state of the website: whether server-side caching is enabled and whether
development (cacheless) mode is on. Works for the main domain, addon websites, parked domains
and subdomains of the account; each reports its own settings and does not inherit them from
its parent website.

`is_enabled` is the setting changed by `Toggle website cache` and reflects a toggle
immediately. `is_cacheless_mode_enabled` is true while development mode is on; caching is then
bypassed even when `is_enabled` is true. `Toggle cacheless mode` is applied in the background,
so this flag can take a few minutes to follow it.

The domain must be the exact website name as listed on the account, without a `www.` prefix.
Returns 404 when the domain is not a website on the account.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting cache status [domain] [flags]
```

### Examples

```
  hostinger hosting cache status mydomain.tld
```

### Options

```
  -h, --help   help for status
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

* [hostinger hosting cache](hostinger_hosting_cache.md)	 - Cache commands

