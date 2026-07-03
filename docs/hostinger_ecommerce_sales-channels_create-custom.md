## hostinger ecommerce sales-channels create-custom

Create custom sales channel

### Synopsis

Create a custom sales channel for a store. Build your own frontend and keep your catalog,
orders, shipping and payments in sync through the Ecommerce API.

```
hostinger ecommerce sales-channels create-custom <store_id> [flags]
```

### Options

```
  -h, --help          help for create-custom
      --name string   Merchant-facing custom name shown in the sales channels list.
      --type string   Sales channel type. Only "custom" channels can be created via the API. (one of: custom)
      --url string    Optional public address where the custom sales channel lives.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce sales-channels](hostinger_ecommerce_sales-channels.md)	 - Sales channels commands

