## hostinger dns records list

Get DNS records

### Synopsis

Retrieve DNS zone records for a specific domain.

Use this endpoint to view current DNS configuration for domain management.

```
hostinger dns records list <domain> [flags]
```

### Examples

```
  hostinger dns records list example.com
  hostinger dns records list example.com --format yaml
```

### Options

```
  -h, --help   help for list
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

* [hostinger dns records](hostinger_dns_records.md)	 - Zone commands

