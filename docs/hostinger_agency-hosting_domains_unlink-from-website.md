## hostinger agency-hosting domains unlink-from-website

Unlink domain from website

### Synopsis

Unlinks a domain from the specified Agency Plan website.

The website stops serving traffic on this domain immediately.

Website files and database are preserved, and any other linked domains remain accessible.

If this is the only domain on the website, unlinking leaves the website without an accessible domain.

```
hostinger agency-hosting domains unlink-from-website <website_uid> <domain> [flags]
```

### Examples

```
  hostinger agency-hosting domains unlink-from-website zpwlGlp19 mydomain.tld
```

### Options

```
  -h, --help   help for unlink-from-website
  -y, --yes    Skip the confirmation prompt
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

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

