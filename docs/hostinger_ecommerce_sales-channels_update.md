## hostinger ecommerce sales-channels update

Update sales channel

### Synopsis

Update a custom sales channel. The merchant-facing `name` and the public `url`
(returned as the channel `domain`) can be changed. Pass `null` to clear a value.

```
hostinger ecommerce sales-channels update <store_id> <sales_channel_id> [flags]
```

### Options

```
  -h, --help          help for update
      --name string   Merchant-facing custom name shown in the sales channels list. Pass null to clear it.
      --url string    Public address where the custom sales channel lives. Pass null to clear it.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce sales-channels](hostinger_ecommerce_sales-channels.md)	 - Sales channels commands

