## hostinger domains transfer list

Get transfer list

### Synopsis

Retrieve all domain transfers in your portfolio.

Use this endpoint to monitor incoming and outgoing registrar transfers across your domains.

```
hostinger domains transfer list [flags]
```

### Examples

```
  hostinger domains transfer list
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

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

