## hapi vps virtual-machines set-root-password

Set root password

### Synopsis

Set root password for a specified virtual machine.

Requirements for password are same as in the [recreate virtual machine
endpoint](/#tag/vps-virtual-machine/POST/api/vps/v1/virtual-machines/{virtualMachineId}/recreate).

Use this endpoint to update administrator credentials for VPS instances.

```
hapi vps virtual-machines set-root-password <virtual-machine-id> [flags]
```

### Options

```
  -h, --help              help for set-root-password
      --password string   Root password for the virtual machine
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

