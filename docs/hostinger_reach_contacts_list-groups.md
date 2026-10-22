## hostinger reach contacts list-groups

List contact groups

### Synopsis

Get a list of all contact groups.

This endpoint returns a list of contact groups that can be used to organize contacts.

```
hostinger reach contacts list-groups [flags]
```

### Examples

```
  hostinger reach contacts list-groups
```

### Options

```
  -h, --help   help for list-groups
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

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

