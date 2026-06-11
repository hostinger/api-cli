## hapi vps ptr delete

Delete PTR record

### Synopsis

This endpoint deletes a PTR (Pointer) record for a specified IP address of a virtual machine.

Once deleted, reverse DNS lookups to the virtual machine's IP address will no longer return the previously configured hostname.

```
hapi vps ptr delete <virtual machine ID> <IP address ID> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps ptr](hapi_vps_ptr.md)	 - PTR records management

