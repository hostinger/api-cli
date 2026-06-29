## hostinger vps ptr delete

Delete PTR record

### Synopsis

Delete a PTR (Pointer) record for a specified virtual machine.

Once deleted, reverse DNS lookups to the virtual machine's IP address will
no longer return the previously configured hostname.

Use this endpoint to remove reverse DNS configuration from VPS instances.

```
hostinger vps ptr delete <virtual-machine-id> <ip-address-id> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps ptr](hostinger_vps_ptr.md)	 - PTR records commands

