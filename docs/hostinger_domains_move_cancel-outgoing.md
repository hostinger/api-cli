## hostinger domains move cancel-outgoing

Cancel outgoing domain move

### Synopsis

Cancel an outgoing move for a specified domain.

The move can only be cancelled while the receiving account has not accepted it yet.
The domain stays in your account.

Use this endpoint to withdraw a move you no longer want to complete.

```
hostinger domains move cancel-outgoing <domain> [flags]
```

### Examples

```
  hostinger domains move cancel-outgoing mydomain.tld
```

### Options

```
  -h, --help   help for cancel-outgoing
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

