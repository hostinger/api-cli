## hostinger vps virtual-machines recreate

Recreate virtual machine

### Synopsis

Recreate a virtual machine from scratch.

The recreation process involves reinstalling the operating system and
resetting the virtual machine to its initial state.
Snapshots, if there are any, will be deleted.

Password Requirements
Password will be checked against leaked password databases.
Requirements for the password are:
- At least 12 characters long
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- Is not leaked publicly

This operation is irreversible and will result in the loss of all data stored on the virtual machine!

Use this endpoint to completely rebuild VPS instances with fresh OS installation.

```
hostinger vps virtual-machines recreate <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines recreate 1268054 --template-id 1130
```

### Options

```
  -h, --help                         help for recreate
      --panel-password string        Panel password for the panel-based OS template. If not provided, random password will be generated. If OS does not support panel_password this field will be ignored. Password will not be shown in the response.
      --password string              Root password for the virtual machine. If not provided, random password will be generated. Password will not be shown in the response.
      --post-install-script-id int   Post-install script to execute after virtual machine was recreated
      --template-id int              Template ID
      --wait                         Wait until the action finishes
      --wait-timeout duration        Maximum time to wait with --wait (default 10m0s)
  -y, --yes                          Skip the confirmation prompt
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

