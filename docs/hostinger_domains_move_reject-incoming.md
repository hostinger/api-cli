## hostinger domains move reject-incoming

Reject incoming domain move

### Synopsis

Reject an incoming move for a specified domain.

The domain stays in the account which initiated the move.
Moves you have already accepted cannot be rejected anymore.

Use this endpoint to decline a domain you do not want to take over.

```
hostinger domains move reject-incoming <domain> [flags]
```

### Examples

```
  hostinger domains move reject-incoming mydomain.tld
```

### Options

```
  -h, --help   help for reject-incoming
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

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

