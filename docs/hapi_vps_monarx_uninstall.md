## hapi vps monarx uninstall

Uninstall Monarx

### Synopsis

Uninstall the Monarx malware scanner on a specified virtual machine.

If Monarx is not installed, the request will still be processed without any effect.

Use this endpoint to remove malware scanner from VPS instances.

```
hapi vps monarx uninstall <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for uninstall
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps monarx](hapi_vps_monarx.md)	 - Malware scanner commands

