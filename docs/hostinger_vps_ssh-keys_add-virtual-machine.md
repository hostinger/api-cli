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

### Options

```
  -h, --help           help for add-virtual-machine
      --keys strings   SSH public keys in OpenSSH format to add
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps ssh-keys](hostinger_vps_ssh-keys.md)	 - SSH Keys commands

