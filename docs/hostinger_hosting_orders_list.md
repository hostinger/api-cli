## hostinger hosting orders list

List orders

### Synopsis

Retrieve a paginated list of orders accessible to the authenticated client.

Only Web and Cloud hosting orders are listed. Agency Plan orders are listed by
`GET /api/agency-hosting/v1/orders`.

This endpoint returns orders of your hosting accounts as well as orders
of other client hosting accounts that have shared access with you.

Use the available query parameters to filter results by order statuses
or specific order IDs for more targeted results.

```
hostinger hosting orders list [flags]
```

### Examples

```
  hostinger hosting orders list
```

### Options

```
      --all                Fetch every page and print all items
  -h, --help               help for list
      --order-ids ints     Filter by specific order IDs
      --page int           Page number
      --per-page int       Number of items per page (default 25)
      --statuses strings   Filter by order statuses (one of: active, deleting, deleted, suspended)
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

* [hostinger hosting orders](hostinger_hosting_orders.md)	 - Orders commands

