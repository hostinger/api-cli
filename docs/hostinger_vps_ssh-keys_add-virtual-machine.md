## hostinger vps ssh-keys add-virtual-machine

Add virtual machine SSH keys

### Synopsis

Add one or more SSH public keys to a specified virtual machine.

Keys are added to the `root` user and can be used for passwordless SSH authentication.
Returns the complete list of SSH keys currently configured on the virtual machine.

Use this endpoint to enable SSH key authentication for VPS instances.

```
hostinger vps ssh-keys add-virtual-machine <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps ssh-keys add-virtual-machine 1268054 --keys 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g... user@example.com'
```

### Options

```
  -h, --help           help for add-virtual-machine
      --keys strings   SSH public keys in OpenSSH format to add
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

* [hostinger vps ssh-keys](hostinger_vps_ssh-keys.md)	 - SSH Keys commands

