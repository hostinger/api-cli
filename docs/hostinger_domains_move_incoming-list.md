## hostinger domains move incoming-list

Get incoming domain move list

### Synopsis

Retrieve all domains other Hostinger accounts are moving to your account.

Moves of every status are returned, including the ones which already completed.

Use this endpoint to find domains waiting for you to accept them.

```
hostinger domains move incoming-list [flags]
```

### Examples

```
  hostinger domains move incoming-list
```

### Options

```
  -h, --help   help for incoming-list
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

