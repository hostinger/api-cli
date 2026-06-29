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

### Options

```
  -h, --help       help for list
      --page int   Page number
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps actions](hostinger_vps_actions.md)	 - Actions commands

