## hapi hosting websites list

List websites

### Synopsis

Retrieve a paginated list of websites (main and addon types) accessible to the authenticated client.

This endpoint returns websites from your hosting accounts as well as
websites from other client hosting accounts that have shared access
with you.

Use the available query parameters to filter results by username,
order ID, enabled status, or domain name for more targeted results.

```
hapi hosting websites list [flags]
```

### Options

```
      --domain string     Filter by domain name (exact match)
  -h, --help              help for list
      --is-enabled        Filter by enabled status
      --order-id int      Order ID
      --page int          Page number
      --per-page int      Number of items per page (default 25)
      --username string   Filter by specific username
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting websites](hapi_hosting_websites.md)	 - Websites commands

