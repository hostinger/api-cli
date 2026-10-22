## hostinger hosting domains list-website-subdomains

List website subdomains

### Synopsis

Retrieve all subdomains created under the selected website.

Use this endpoint to inspect subdomain configuration for a specific website,
including the parent domain and root directory assigned to each subdomain.

```
hostinger hosting domains list-website-subdomains <username> <domain> [flags]
```

### Examples

```
  hostinger hosting domains list-website-subdomains u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for list-website-subdomains
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

