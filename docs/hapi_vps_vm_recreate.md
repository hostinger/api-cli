## hapi vps vm recreate

Recreate virtual machine

### Synopsis

This endpoint will recreate a virtual machine from scratch. The recreation process involves reinstalling the 
operating system and resetting the virtual machine to its initial state. Snapshots, if there are any, will be deleted.

```
hapi vps vm recreate <virtual machine ID> [flags]
```

### Options

```
  -h, --help                         help for recreate
      --password string              Root or panel Password
      --post_install_script_id int   Post install script (default -1)
      --template_id int              Template ID (default -1)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

