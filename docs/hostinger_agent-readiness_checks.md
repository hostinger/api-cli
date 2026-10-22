## hostinger agent-readiness checks

Checks commands

```
hostinger agent-readiness checks [flags]
```

### Options

```
  -h, --help   help for checks
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

* [hostinger agent-readiness](hostinger_agent-readiness.md)	 - Agent Readiness commands
* [hostinger agent-readiness checks list](hostinger_agent-readiness_checks_list.md)	 - List checks

