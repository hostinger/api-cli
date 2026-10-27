## hostinger vps snapshots create

Create snapshot

### Synopsis

Create a snapshot of a specified virtual machine.

A snapshot captures the state and data of the virtual machine at a specific point in time,
allowing users to restore the virtual machine to that state if needed.
This operation is useful for backup purposes, system recovery,
and testing changes without affecting the current state of the virtual machine.

Creating new snapshot will overwrite the existing snapshot!

Use this endpoint to capture VPS state for backup and recovery purposes.

```
hostinger vps snapshots create <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps snapshots create 1268054
```

### Options

```
  -h, --help                    help for create
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps snapshots](hostinger_vps_snapshots.md)	 - Snapshots commands

