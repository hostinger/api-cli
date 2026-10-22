## hostinger domains transfer get

Get transfer

### Synopsis

Retrieve the transfer for a specified domain.

Use this endpoint to track an incoming or outgoing registrar transfer and its status.

```
hostinger domains transfer get <domain> [flags]
```

### Examples

```
  hostinger domains transfer get mydomain.tld
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

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

