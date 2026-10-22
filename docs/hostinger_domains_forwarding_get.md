## hostinger domains forwarding get

Get domain forwarding

### Synopsis

Retrieve domain forwarding data.

Use this endpoint to view current redirect configuration for domains.

```
hostinger domains forwarding get <domain> [flags]
```

### Examples

```
  hostinger domains forwarding get mydomain.tld
```

### Options

```
  -h, --help   help for get
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

