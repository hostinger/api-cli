## hapi vps virtual-machines set-hostname

Set hostname

### Synopsis

Set hostname for a specified virtual machine.

Changing hostname does not update PTR record automatically.
If you want your virtual machine to be reachable by a hostname, 
you need to point your domain A/AAAA records to virtual machine IP as well.

Use this endpoint to configure custom hostnames for VPS instances.

```
hapi vps virtual-machines set-hostname <virtual-machine-id> [flags]
```

### Options

```
  -h, --help              help for set-hostname
      --hostname string   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

