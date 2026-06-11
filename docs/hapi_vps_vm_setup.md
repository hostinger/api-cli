## hapi vps vm setup

Setup purchased virtual machine

### Synopsis

This endpoint sets up a newly purchased virtual machine that has not been set up yet, by installing the
selected operating system template in the chosen data center.

```
hapi vps vm setup <virtual machine ID> [flags]
```

### Options

```
      --data_center_id int           Data center ID (default -1)
      --enable_backups               Enable weekly backup schedule
  -h, --help                         help for setup
      --hostname string              Override default hostname of the virtual machine
      --install_monarx               Install Monarx malware scanner (if supported)
      --ns1 string                   Name server 1
      --ns2 string                   Name server 2
      --password string              Password for the virtual machine. If not provided, a random password will be generated
      --post_install_script_id int   Post-install script ID (default -1)
      --public_key string            Contents of the SSH key to use
      --public_key_name string       Name of the SSH key to use
      --template_id int              Template ID (default -1)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps vm](hapi_vps_vm.md)	 - Virtual machine management

