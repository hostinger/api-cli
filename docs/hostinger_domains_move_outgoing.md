## hostinger domains move outgoing

Get outgoing domain move

### Synopsis

Retrieve the outgoing move for a specified domain.

Returns 404 when the domain has no move in progress.

Use this endpoint to track the status of a move you have initiated for a single domain.

```
hostinger domains move outgoing <domain> [flags]
```

### Examples

```
  hostinger domains move outgoing mydomain.tld
```

### Options

```
  -h, --help   help for outgoing
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

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

