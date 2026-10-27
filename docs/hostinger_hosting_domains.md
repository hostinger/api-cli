## hostinger hosting domains

Domains commands

```
hostinger hosting domains [flags]
```

### Options

```
  -h, --help   help for domains
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

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting domains create-website-parked](hostinger_hosting_domains_create-website-parked.md)	 - Create website parked domain
* [hostinger hosting domains create-website-subdomain](hostinger_hosting_domains_create-website-subdomain.md)	 - Create website subdomain
* [hostinger hosting domains delete-website-parked](hostinger_hosting_domains_delete-website-parked.md)	 - Delete website parked domain
* [hostinger hosting domains delete-website-subdomain](hostinger_hosting_domains_delete-website-subdomain.md)	 - Delete website subdomain
* [hostinger hosting domains generate-free-subdomain](hostinger_hosting_domains_generate-free-subdomain.md)	 - Generate a free subdomain
* [hostinger hosting domains list-website-parked](hostinger_hosting_domains_list-website-parked.md)	 - List website parked domains
* [hostinger hosting domains list-website-subdomains](hostinger_hosting_domains_list-website-subdomains.md)	 - List website subdomains
* [hostinger hosting domains verify-ownership](hostinger_hosting_domains_verify-ownership.md)	 - Verify domain ownership

