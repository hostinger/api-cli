## hostinger hosting cache toggle-website

Toggle website cache

### Synopsis

Turns server-side caching for the website on or off, based on the enabled flag. Enable it for
faster page loads, reduced server load, and improved user experience; recommended for production
websites. Disabling may impact performance; to temporarily bypass caching while developing or
debugging, prefer toggling cacheless mode instead.

Does nothing if caching is already in the requested state.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting cache toggle-website [domain] [flags]
```

### Examples

```
  hostinger hosting cache toggle-website mydomain.tld --enabled=true
```

### Options

```
      --enabled   Turn server-side caching on (true) or off (false) for the website.
  -h, --help      help for toggle-website
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

