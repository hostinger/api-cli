## hostinger vps virtual-machines set-panel-password

Set panel password

### Synopsis

Set panel password for a specified virtual machine.

If virtual machine does not use panel OS, the request will still be processed without any effect.
Requirements for password are same as in the recreate virtual machine
endpoint (/#tag/vps-virtual-machine/POST/api/vps/v1/virtual-machines/{virtualMachineId}/recreate).

Use this endpoint to configure control panel access credentials for VPS instances.

```
hostinger vps virtual-machines set-panel-password <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines set-panel-password 1268054 --password oMeNRustosIO
```

### Options

```
  -h, --help                    help for set-panel-password
      --password string         Panel password for the virtual machine
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

