## hostinger reach automations list-steps

List automation steps

### Synopsis

Get the workflow of an automation as a flat list of steps.

The steps form a tree rather than a straight line: follow `parent_uuid` to reconstruct the
branches, and use `step_order` to order the steps that share a parent. An automation with no
steps yet returns an empty list.

```
hostinger reach automations list-steps <profile-uuid> <automation-uuid> [flags]
```

### Options

```
  -h, --help   help for list-steps
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach automations](hostinger_reach_automations.md)	 - Automations commands

