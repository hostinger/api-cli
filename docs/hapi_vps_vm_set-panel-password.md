## hapi vps vm set-panel-password

Set panel password

### Synopsis

This endpoint sets the panel password for a specified virtual machine. 
If virtual machine does not use panel OS, the request will still be processed without any effect.

```
hapi vps vm set-panel-password <virtual machine ID> [flags]
```

### Options

```
  -h, --help   help for set-panel-password
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

