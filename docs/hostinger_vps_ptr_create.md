## hostinger vps ptr create

Create PTR record

### Synopsis

Create or update a PTR (Pointer) record for a specified virtual machine.

Use this endpoint to configure reverse DNS lookup for VPS IP addresses.

```
hostinger vps ptr create <virtual-machine-id> <ip-address-id> [flags]
```

### Options

```
      --domain string   Pointer record domain
  -h, --help            help for create
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps ptr](hostinger_vps_ptr.md)	 - PTR records commands

