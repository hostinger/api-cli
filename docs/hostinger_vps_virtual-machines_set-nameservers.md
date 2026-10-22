## hostinger vps virtual-machines set-nameservers

Set nameservers

### Synopsis

Set nameservers for a specified virtual machine.

Be aware, that improper nameserver configuration can lead to the virtual
machine being unable to resolve domain names.

Use this endpoint to configure custom DNS resolvers for VPS instances.

```
hostinger vps virtual-machines set-nameservers <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines set-nameservers 1268054 --ns1 4.3.2.1
```

### Options

```
  -h, --help                    help for set-nameservers
      --ns1 string              
      --ns2 string              
      --ns3 string              
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
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

