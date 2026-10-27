## hostinger hosting cache clear-website

Clear website cache

### Synopsis

Permanently clears all server-side cache for the website at once. Use it when content was
updated and needs to be visible immediately, or after making major changes.

Also purges the Hostinger CDN cache when CDN is enabled on the website. To purge only the CDN
cache and leave the server cache untouched, use `Purge CDN cache`. For a WordPress
installation living in a subdirectory, pass the directory query parameter to clear its cache.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting cache clear-website [domain] [flags]
```

### Examples

```
  hostinger hosting cache clear-website mydomain.tld
```

### Options

```
      --directory string   Directory of the website installation to clear, relative to the website root. Defaults to the website root.
  -h, --help               help for clear-website
  -y, --yes                Skip the confirmation prompt
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

