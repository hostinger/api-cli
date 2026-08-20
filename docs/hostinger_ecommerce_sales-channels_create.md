## hostinger ecommerce sales-channels create

Create a sales channel

### Synopsis

Create a sales channel for a store. A "custom" channel is headless: build your own frontend and keep
your catalog, orders, shipping and payments in sync through the Ecommerce API. A "quick-link" channel
is a hosted one-page store whose handle is auto-generated.

```
hostinger ecommerce sales-channels create <store_id> [flags]
```

### Options

```
  -h, --help          help for create
      --name string   Merchant-facing custom name. Required for custom channels; not supported for quick-link.
      --type string   Sales channel type. "custom" is a headless channel: it requires a name and takes an optional public url.
                      "quick-link" is a one-page store whose handle is auto-generated; it supports neither name nor url. (one of: custom, quick-link)
      --url string    Optional public url for the channel. Custom channels only; not supported for quick-link.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger ecommerce sales-channels](hostinger_ecommerce_sales-channels.md)	 - Sales channels commands

