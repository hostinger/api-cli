## hostinger domains forwarding delete

Delete domain forwarding

### Synopsis

Delete domain forwarding data.

Use this endpoint to remove redirect configuration from domains.

```
hostinger domains forwarding delete <domain> [flags]
```

### Examples

```
  hostinger domains forwarding delete mydomain.tld
```

### Options

```
  -h, --help   help for delete
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

* [hostinger domains forwarding](hostinger_domains_forwarding.md)	 - Forwarding commands

