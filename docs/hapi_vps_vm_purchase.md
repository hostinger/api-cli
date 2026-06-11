## hapi vps vm purchase

Purchase and setup a new virtual machine

### Synopsis

This endpoint purchases a new virtual machine using the provided catalog price item and payment method, and
sets it up with the selected operating system template in the chosen data center.

The purchased virtual machine will be set for automatic renewal.

```
hapi vps vm purchase [flags]
```

### Options

```
  -c, --coupon stringArray           Discount coupon code (repeatable)
      --data_center_id int           Data center ID (default -1)
      --enable_backups               Enable weekly backup schedule
  -h, --help                         help for purchase
      --hostname string              Override default hostname of the virtual machine
      --install_monarx               Install Monarx malware scanner (if supported)
      --item_id string               Catalog price item ID
      --ns1 string                   Name server 1
      --ns2 string                   Name server 2
      --password string              Password for the virtual machine. If not provided, a random password will be generated
      --payment_method_id int        Payment method ID (default payment method is used if not provided) (default -1)
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

