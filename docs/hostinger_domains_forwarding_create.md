## hostinger domains forwarding create

Create domain forwarding

### Synopsis

Create domain forwarding configuration.

Creating a forwarding changes the domain's DNS. If the domain does not use Hostinger nameservers, they are
switched to Hostinger's nameservers automatically, so website, email and other services whose DNS is hosted
elsewhere stop working. The root (@) A and AAAA records in the domain's Hostinger DNS zone are also replaced
with records pointing to Hostinger, even when the domain already uses Hostinger nameservers. The domain's
current nameservers are returned by the domain details endpoint (#tag/domains-portfolio); confirm with the
user before creating a forwarding on a domain that uses other nameservers. DNS changes can take some time to
propagate.

Use this endpoint to set up domain redirects to other URLs.

```
hostinger domains forwarding create [flags]
```

### Examples

```
  hostinger domains forwarding create --domain mydomain.tld --redirect-type 301 --redirect-url https://forward.to.my.url
```

### Options

```
      --domain string          Domain name
  -h, --help[=true|false]      help for create
      --redirect-type string   Redirect type (one of: 301, 302)
      --redirect-url string    URL to forward domain to
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger domains forwarding](hostinger_domains_forwarding.md)	 - Forwarding commands

