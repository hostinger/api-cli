## hostinger ecommerce sales-channels update

Update sales channel

### Synopsis

Update a custom sales channel. The merchant-facing `name` and the public `url`
(returned as the channel `domain`) can be changed. Pass `null` to clear a value.

```
hostinger ecommerce sales-channels update <store_id> <sales_channel_id> [flags]
```

### Examples

```
  hostinger ecommerce sales-channels update store_01J8Z5F8W9K8M4A7B3C2D1E0FG scha_01J8Z5F8W9K8M4A7B3C2D1E0FG
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce sales-channels](hostinger_ecommerce_sales-channels.md)	 - Sales channels commands

