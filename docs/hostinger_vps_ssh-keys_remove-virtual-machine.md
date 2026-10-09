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

### Options

```
  -h, --help           help for remove-virtual-machine
      --keys strings   SSH public keys in OpenSSH format to remove
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps ssh-keys](hostinger_vps_ssh-keys.md)	 - SSH Keys commands

