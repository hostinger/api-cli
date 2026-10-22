## hostinger vps actions list

Get actions

### Synopsis

Retrieve actions performed on a specified virtual machine.

Actions are operations or events that have been executed on the virtual
machine, such as starting, stopping, or modifying the machine. This endpoint
allows you to view the history of these actions, providing details about
each action, such as the action name, timestamp, and status.

Use this endpoint to view VPS operation history and troubleshoot issues.

```
hostinger vps actions list <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps actions list 1268054
```

### Options

```
      --all        Fetch every page and print all items
  -h, --help       help for list
      --page int   Page number
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

* [hostinger vps actions](hostinger_vps_actions.md)	 - Actions commands

