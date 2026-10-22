## hostinger domains move outgoing-list

Get outgoing domain move list

### Synopsis

Retrieve all domains you are moving to other Hostinger accounts.

Only moves which have not completed yet are returned.

Use this endpoint to track moves you have initiated and the accounts they are addressed to.

```
hostinger domains move outgoing-list [flags]
```

### Examples

```
  hostinger domains move outgoing-list
```

### Options

```
  -h, --help   help for outgoing-list
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

