## hostinger ecommerce sales-channels create

Create a sales channel

### Synopsis

Create a sales channel for a store. A "custom" channel is headless: build your own frontend and keep
your catalog, orders, shipping and payments in sync through the Ecommerce API. A "quick-link" channel
is a hosted one-page store whose handle is auto-generated.

```
hostinger ecommerce sales-channels create <store_id> [flags]
```

### Examples

```
  hostinger ecommerce sales-channels create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --type custom
```

### Options

```
  -h, --help          help for create
      --name string   Merchant-facing custom name. Required for custom channels; not supported for quick-link.
      --type string   Sales channel type. "custom" is a headless channel: it requires a name and takes an optional public url. "quick-link" is a one-page store whose handle is auto-generated; it supports neither name nor url. (one of: custom, quick-link)
      --url string    Optional public url for the channel. Custom channels only; not supported for quick-link.
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce sales-channels](hostinger_ecommerce_sales-channels.md)	 - Sales channels commands

