## hostinger vps ssh-keys remove-virtual-machine

Remove virtual machine SSH keys

### Synopsis

Remove one or more SSH public keys from a specified virtual machine.

Removed keys can no longer be used to authenticate via SSH as the `root` user.
Returns the remaining list of SSH keys configured on the virtual machine.

Use this endpoint to revoke SSH key access to VPS instances.

```
hostinger vps ssh-keys remove-virtual-machine <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps ssh-keys remove-virtual-machine 1268054 --keys 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g... user@example.com'
```

### Options

```
  -h, --help           help for remove-virtual-machine
      --keys strings   SSH public keys in OpenSSH format to remove
  -y, --yes            Skip the confirmation prompt
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

