## hostinger vps public-keys list

Get public keys

### Synopsis

Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per
virtual machine via `GET /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.

Retrieve public keys associated with your account.

Use this endpoint to view available SSH keys for VPS authentication.

```
hostinger vps public-keys list [flags]
```

### Examples

```
  hostinger vps public-keys list
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

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

