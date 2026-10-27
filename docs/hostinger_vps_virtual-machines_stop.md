## hostinger vps virtual-machines stop

Stop virtual machine

### Synopsis

Stop a specified virtual machine.

If the virtual machine is already stopped, the request will still be processed without any effect.

This is a compute-only power state change and does not affect billing. To stop future charges,
disable auto-renewal on the owning subscription.

Use this endpoint to power off running VPS instances.

```
hostinger vps virtual-machines stop <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines stop 1268054
```

### Options

```
  -h, --help                    help for stop
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

