## hapi vps virtual-machines setup

Setup purchased virtual machine

### Synopsis

Setup newly purchased virtual machine with `initial` state.

Use this endpoint to configure and initialize purchased VPS instances.

```
hapi vps virtual-machines setup <virtual-machine-id> [flags]
```

### Options

```
      --data-center-id int           Data center ID
      --enable-backups               Enable weekly backup schedule (default true)
  -h, --help                         help for setup
      --hostname string              Override default hostname of the virtual machine
      --install-monarx               Install Monarx malware scanner (if supported)
      --ns1 string                   Name server 1
      --ns2 string                   Name server 2
      --password string              Password for the virtual machine. If not provided, random password will be generated.
                                     Password will not be shown in the response.
      --post-install-script-id int   Post-install script ID
      --public-key string            Use SSH key (JSON)
      --template-id int              Template ID
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps virtual-machines](hapi_vps_virtual-machines.md)	 - Virtual machine commands

