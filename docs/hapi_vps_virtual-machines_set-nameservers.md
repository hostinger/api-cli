## hapi vps virtual-machines set-nameservers

Set nameservers

### Synopsis

Set nameservers for a specified virtual machine.

Be aware, that improper nameserver configuration can lead to the virtual
machine being unable to resolve domain names.

Use this endpoint to configure custom DNS resolvers for VPS instances.

```
hapi vps virtual-machines set-nameservers <virtual-machine-id> [flags]
```

### Options

```
  -h, --help         help for set-nameservers
      --ns1 string   
      --ns2 string   
      --ns3 string   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

