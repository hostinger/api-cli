## hostinger vps virtual-machines purchase

Purchase new virtual machine

### Synopsis

Purchase and setup a new virtual machine.

If virtual machine setup fails for any reason, login to
hPanel (https://hpanel.hostinger.com/) and complete the setup manually.

If no payment method is provided, your default payment method will be used automatically.

If the response is `202 Accepted`, the payment is still being processed and the virtual machine
was not set up. Login to
hPanel (https://hpanel.hostinger.com/) and complete the setup manually.

Use this endpoint to create new VPS instances.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --coupons
    [
      "Coupon 3"
    ]

  --setup
    {
      "data_center_id": 19,
      "enable_backups": true,
      "hostname": "my.server.tld",
      "install_monarx": false,
      "ns1": "4.3.2.1",
      "ns2": "1.2.3.4",
      "password": "oMeNRustosIO",
      "post_install_script_id": 6324,
      "public_key": {
        "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2X...",
        "name": "my-key"
      },
      "template_id": 1130
    }

```
hostinger vps virtual-machines purchase [flags]
```

### Examples

```
  hostinger vps virtual-machines purchase --item-id hostingercom-vps-kvm2-usd-1m --setup @setup.json
```

### Options

```
      --coupons string          Discount coupon codes (JSON, @file or -; shape in the help above)
  -h, --help                    help for purchase
      --item-id string          Catalog price item ID
      --payment-method-id int   Payment method ID, default will be used if not provided
      --setup string            (JSON, @file or -; shape in the help above)
  -y, --yes                     Skip the confirmation prompt
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

