## hostinger hosting domains generate-free-subdomain

Generate a free subdomain

### Synopsis

Generate a unique free subdomain that can be used for hosting services without purchasing custom domains.
Free subdomains allow you to start using hosting services immediately
and you can always connect a custom domain to your site later.

```
hostinger hosting domains generate-free-subdomain [flags]
```

### Examples

```
  hostinger hosting domains generate-free-subdomain
```

### Options

```
  -h, --help   help for generate-free-subdomain
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

* [hostinger hosting domains](hostinger_hosting_domains.md)	 - Domains commands

