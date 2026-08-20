## hostinger hosting websites list

List websites

### Synopsis

Retrieve a paginated list of websites (CloudLinux, Builder, and Horizons) accessible to the
authenticated client.

This endpoint returns websites from your hosting accounts as well as
websites from other client hosting accounts that have shared access
with you.

Each website includes a `website_type` field describing the type of
website detected on the underlying platform (`wordpress`, `builder`,
`horizons`, `nodejs`, or `other`). Some fields, such as
`vhost_type`, `username`, and `root_directory`, only apply to
CloudLinux websites and are null for other platforms.

Use `website_types` to list only websites of a given detected type, e.g. only
WordPress websites (`website_types=wordpress`) or only Node.js websites
(`website_types=nodejs`). Combine with the other available query parameters to
filter by username, order ID, enabled status, or domain name for more targeted
results.

```
hostinger hosting websites list [flags]
```

### Options

```
      --domain string           Filter by domain name (case-insensitive substring match)
  -h, --help                    help for list
      --is-enabled              Filter by enabled status
      --order-id int            Order ID
      --page int                Page number
      --per-page int            Number of items per page (default 25)
      --username string         Filter by specific username
      --website-types strings   Filter by detected website type, e.g. wordpress,nodejs. Accepts a comma-separated list. (one of: wordpress, builder, horizons, nodejs, other)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

