## hapi vps vm restart

Restart virtual machine

### Synopsis

This endpoint restarts a specified virtual machine. This is equivalent to fully stopping and starting the virtual machine.

If the virtual machine was stopped, then it will be started.

```
hapi vps vm restart <virtual machine ID> [flags]
```

### Options

```
  -h, --help   help for restart
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

