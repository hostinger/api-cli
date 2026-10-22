## hostinger hosting domains create-website-parked

Create website parked domain

### Synopsis

Create a parked or alias domain for the selected website.

Provide a domain name or IP address to park on the website so it serves the same content
as the parent domain.

```
hostinger hosting domains create-website-parked <username> <domain> [flags]
```

### Examples

```
  hostinger hosting domains create-website-parked u123456789 mydomain.tld --parked-domain parked-domain.com
```

### Options

```
  -h, --help                   help for create-website-parked
      --parked-domain string   Domain name or IP address to park on the selected website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting domains](hostinger_hosting_domains.md)	 - Domains commands

