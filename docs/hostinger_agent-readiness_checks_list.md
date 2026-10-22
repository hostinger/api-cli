## hostinger agent-readiness checks list

List checks

### Synopsis

List every check an agent readiness scan runs, with its name, category and weight.\n
Use it to show the human-readable name of each `check_id` in a scan's results.

```
hostinger agent-readiness checks list [flags]
```

### Examples

```
  hostinger agent-readiness checks list
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

* [hostinger agent-readiness checks](hostinger_agent-readiness_checks.md)	 - Checks commands

