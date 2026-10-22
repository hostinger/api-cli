## hostinger hosting domains delete-website-parked

Delete website parked domain

### Synopsis

Delete an existing parked or alias domain from the selected website.

Use this endpoint to remove parked domains that are no longer needed.

```
hostinger hosting domains delete-website-parked <username> <domain> <parked-domain> [flags]
```

### Examples

```
  hostinger hosting domains delete-website-parked u123456789 mydomain.tld parked-domain.com
```

### Options

```
  -h, --help   help for delete-website-parked
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

