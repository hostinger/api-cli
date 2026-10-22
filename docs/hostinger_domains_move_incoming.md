## hostinger domains move incoming

Get incoming domain move

### Synopsis

Retrieve the incoming move for a specified domain.

Returns 404 when no account is moving this domain to you.

Use this endpoint to check whether a domain addressed to you is still waiting to be accepted.

```
hostinger domains move incoming <domain> [flags]
```

### Examples

```
  hostinger domains move incoming mydomain.tld
```

### Options

```
      --force-sync activating   Re-check the move against the registry before responding. Only has an effect while the move is in the activating status.
  -h, --help                    help for incoming
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

