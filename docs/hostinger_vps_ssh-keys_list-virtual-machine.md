## hostinger vps ssh-keys list-virtual-machine

List virtual machine SSH keys

### Synopsis

Retrieve SSH public keys currently configured on a specified virtual machine.

Only keys of the `root` user are listed.

Use this endpoint to view SSH keys that can be used for authentication on VPS instances.

```
hostinger vps ssh-keys list-virtual-machine <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps ssh-keys list-virtual-machine 1268054
```

### Options

```
  -h, --help   help for list-virtual-machine
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

