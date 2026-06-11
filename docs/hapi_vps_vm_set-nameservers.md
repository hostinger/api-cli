## hapi vps vm set-nameservers

Set nameservers

### Synopsis

This endpoint sets the nameservers for a specified virtual machine. 
Be aware, that improper nameserver configuration can lead to the virtual machine being unable to resolve domain names.

```
hapi vps vm set-nameservers <virtual machine ID> [flags]
```

### Options

```
  -h, --help   help for set-nameservers
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

