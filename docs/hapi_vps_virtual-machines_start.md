## hapi vps virtual-machines start

Start virtual machine

### Synopsis

Start a specified virtual machine.

If the virtual machine is already running, the request will still be processed without any effect.

Use this endpoint to power on stopped VPS instances.

```
hapi vps virtual-machines start <virtual-machine-id> [flags]
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

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

