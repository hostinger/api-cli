## hostinger agency-hosting domains link-to-website

Link domain to website

### Synopsis

Links a domain to the specified Agency Plan website so it can serve traffic for that domain.

```
hostinger agency-hosting domains link-to-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting domains link-to-website zpwlGlp19 --domain example.com
```

### Options

```
      --domain string   Fully qualified domain name to link to the website
  -h, --help            help for link-to-website
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

