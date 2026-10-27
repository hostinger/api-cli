## hostinger vps backups restore

Restore backup

### Synopsis

Restore a backup for a specified virtual machine.

The system will then initiate the restore process, which may take some time depending on the size of the backup.

All data on the virtual machine will be overwritten with the data from the backup.

Use this endpoint to recover VPS data from backup points.

```
hostinger vps backups restore <virtual-machine-id> <backup-id> [flags]
```

### Examples

```
  hostinger vps backups restore 1268054 8676502
```

### Options

```
  -h, --help                    help for restore
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
  -y, --yes                     Skip the confirmation prompt
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

* [hostinger vps backups](hostinger_vps_backups.md)	 - Backups commands

