## hostinger vps virtual-machines purchase

Purchase new virtual machine

### Synopsis

Purchase and setup a new virtual machine.

If virtual machine setup fails for any reason, login to
[hPanel](https://hpanel.hostinger.com/) and complete the setup manually.

If no payment method is provided, your default payment method will be used automatically.

Use this endpoint to create new VPS instances.

```
hostinger vps virtual-machines purchase [flags]
```

### Options

```
      --coupons string          Discount coupon codes (JSON)
  -h, --help                    help for purchase
      --item-id string          Catalog price item ID
      --payment-method-id int   Payment method ID, default will be used if not provided
      --setup string             (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

