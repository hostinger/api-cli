## hostinger hosting domains delete-website-subdomain

Delete website subdomain

### Synopsis

Delete an existing subdomain from the selected website.

Use this endpoint to remove subdomains that are no longer needed.

```
hostinger hosting domains delete-website-subdomain <username> <domain> <subdomain> [flags]
```

### Examples

```
  hostinger hosting domains delete-website-subdomain u123456789 mydomain.tld blog
```

### Options

```
  -h, --help   help for delete-website-subdomain
  -y, --yes    Skip the confirmation prompt
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

