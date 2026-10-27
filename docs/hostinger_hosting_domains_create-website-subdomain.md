## hostinger hosting domains create-website-subdomain

Create website subdomain

### Synopsis

Create a new subdomain for the selected website.

Provide a subdomain prefix and, optionally, a custom directory or the
website public directory to use as the subdomain root.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting domains create-website-subdomain [domain] [flags]
```

### Examples

```
  hostinger hosting domains create-website-subdomain mydomain.tld --subdomain blog
```

### Options

```
      --directory string            Directory name for the subdomain relative to the website root
  -h, --help                        help for create-website-subdomain
      --is-using-public-directory   Use the website public directory as the subdomain root directory
      --subdomain string            Subdomain prefix to create under the selected website
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

