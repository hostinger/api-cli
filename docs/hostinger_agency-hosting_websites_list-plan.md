## hostinger agency-hosting websites list-plan

List Agency Plan websites

### Synopsis

Retrieve a paginated list of Agency Plan websites (H5G, Builder, and Horizons) accessible to
the authenticated client.

This endpoint returns websites from your hosting accounts as well as
websites from other client hosting accounts that have shared access
with you.

The response shape differs per platform — see the `platform` field on each item.

Use `website_types` to list only websites of a given detected type, e.g. only
WordPress websites (`website_types=wordpress`) or only Node.js websites
(`website_types=nodejs`). Combine with `order_ids`, `states`, or `domain` for more
targeted results.

```
hostinger agency-hosting websites list-plan [flags]
```

### Examples

```
  hostinger agency-hosting websites list-plan
```

### Options

```
      --all                     Fetch every page and print all items
      --domain string           Filter by domain name (case-insensitive substring match)
  -h, --help                    help for list-plan
      --order-ids ints          Filter by order IDs. Accepts a comma-separated list.
      --page int                Page number
      --per-page int            Number of items per page (default 25)
      --states strings          Filter by website state. Accepts a comma-separated list. (one of: active, locked, suspended, deleting, deleted)
      --website-types strings   Filter by detected website type, e.g. wordpress,nodejs. Accepts a comma-separated list. (one of: wordpress, builder, horizons, nodejs, other)
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

* [hostinger agency-hosting websites](hostinger_agency-hosting_websites.md)	 - Websites commands

