## hapi vps vm start

Start a virtual machine

### Synopsis

This endpoint starts a specified virtual machine. If the virtual machine is already running, 
the request will still be processed without any effect.

```
hapi vps vm start <virtual machine ID> [flags]
```

### Options

```
  -h, --help   help for start
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

