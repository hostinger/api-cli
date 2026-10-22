## hostinger vps virtual-machines attached-public-keys

Get attached public keys

### Synopsis

Deprecated: use `GET /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys` instead.

Retrieve public keys attached to a specified virtual machine.

Use this endpoint to view SSH keys configured for specific VPS instances.

```
hostinger vps virtual-machines attached-public-keys <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines attached-public-keys 1268054
```

### Options

```
      --all        Fetch every page and print all items
  -h, --help       help for attached-public-keys
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

