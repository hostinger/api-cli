## hostinger reach automations get

Get automation details

### Synopsis

Get a single automation with the counts of contacts that entered it, are moving through it,
finished it or failed on the way.

This describes the automation itself. To see the workflow it runs, use the steps endpoint.

```
hostinger reach automations get <profile-uuid> <automation-uuid> [flags]
```

### Examples

```
  hostinger reach automations get 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
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

* [hostinger reach automations](hostinger_reach_automations.md)	 - Automations commands

