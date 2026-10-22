## hostinger ecommerce orders list-store

List store orders

### Synopsis

List a store's orders newest first as summaries. Filter by status, payment or fulfilment
status, customer email, order number or a free-text query. Amounts are in the smallest
currency unit. Retrieve a single order for its line items, addresses and fulfilments.

```
hostinger ecommerce orders list-store <store_id> [flags]
```

### Examples

```
  hostinger ecommerce orders list-store store_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
      --all                          Fetch every page and print all items
      --created-at-from string       Earliest creation time to include, inclusive. Accepts a date or ISO date-time (UTC).
      --created-at-to string         Latest creation time to include, inclusive. A bare date covers that whole day.
      --display-id string            The order number the merchant and customer see.
      --email string                 Customer email, matched exactly.
      --fulfillment-status strings   Fulfilment statuses to include. (one of: not_fulfilled, partially_fulfilled, fulfilled, partially_shipped, shipped, partially_returned, returned, canceled, requires_action)
  -h, --help                         help for list-store
      --page int                     Page number
      --payment-status strings       Payment statuses to include. A paid order is "captured". (one of: not_paid, awaiting, captured, partially_refunded, refunded, canceled, requires_action, not_required)
      --q string                     Free-text search over customer name, email, order number and line items.
      --status strings               Order statuses to include. (one of: pending, completed, archived, canceled, requires_action)
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

* [hostinger ecommerce orders](hostinger_ecommerce_orders.md)	 - Orders commands

