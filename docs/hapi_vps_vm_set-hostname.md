## hapi vps vm set-hostname

Set hostname

### Synopsis

This endpoint sets the hostname for a specified virtual machine. Changing hostname does not update PTR record automatically.
If you want your virtual machine to be reachable by a hostname, you need to point your domain A/AAAA records 
to virtual machine IP as well.

```
hapi vps vm set-hostname <virtual machine ID> [flags]
```

### Options

```
  -h, --help              help for set-hostname
      --hostname string   Virtual machine hostname
      --ns1 string        Name server 1
      --ns2 string        Name server 2
      --password string   Password
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

