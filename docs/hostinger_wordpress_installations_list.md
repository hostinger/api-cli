## hostinger wordpress installations list

List WordPress installations

### Synopsis

List WordPress installations on shared and cloud web hosting accounts (the
`/api/hosting/v1` family, identified by `username`) accessible to the
authenticated client.

Agency Plan websites are not included. To work with WordPress on an Agency Plan
website, list them with GET /api/agency-hosting/v1/websites?website_types=wordpress
and read their details with
GET /api/agency-hosting/v1/websites/{website_uid}/wordpress/settings. An account
that only has Agency Plan websites receives an empty list from this endpoint.

Use this endpoint to discover existing WordPress installations and to poll
for installation status after calling the install endpoint. When a newly
requested installation appears in this list, WordPress is ready. Filter by
username and domain to narrow results to a specific website.

Each installation includes a `valid` flag and, when invalid, a
`validationError` describing why.

```
hostinger wordpress installations list [flags]
```

### Examples

```
  hostinger wordpress installations list
```

### Options

```
      --domain string      Filter by domain name (case-insensitive substring match)
  -h, --help               help for list
      --ownership string   Filter by ownership type. Defaults to "owned". Use "all" to include both owned and managed installations. (one of: owned, managed, all)
      --username string    Filter by specific username
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

