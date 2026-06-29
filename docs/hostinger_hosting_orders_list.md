## hostinger hosting orders list

List orders

### Synopsis

Retrieve a paginated list of orders accessible to the authenticated client.

This endpoint returns orders of your hosting accounts as well as orders
of other client hosting accounts that have shared access with you.

Use the available query parameters to filter results by order statuses
or specific order IDs for more targeted results.

```
hostinger hosting orders list [flags]
```

### Options

```
  -h, --help               help for list
      --order-ids ints     Filter by specific order IDs
      --page int           Page number
      --per-page int       Number of items per page (default 25)
      --statuses strings   Filter by order statuses (one of: active, deleting, deleted, suspended)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting orders](hostinger_hosting_orders.md)	 - Orders commands

