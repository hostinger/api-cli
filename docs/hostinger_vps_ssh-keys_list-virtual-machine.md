## hostinger vps ssh-keys list-virtual-machine

List virtual machine SSH keys

### Synopsis

Retrieve SSH public keys currently configured on a specified virtual machine.

Only keys of the `root` user are listed.

Use this endpoint to view SSH keys that can be used for authentication on VPS instances.

```
hostinger vps ssh-keys list-virtual-machine <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for list-virtual-machine
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps ssh-keys](hostinger_vps_ssh-keys.md)	 - SSH Keys commands

