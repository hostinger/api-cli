## hostinger wordpress installations list

List WordPress installations

### Synopsis

List WordPress installations accessible to the authenticated client.

Use this endpoint to discover existing WordPress installations and to poll
for installation status after calling the install endpoint. When a newly
requested installation appears in this list, WordPress is ready. Filter by
username and domain to narrow results to a specific website.

Each installation includes a `valid` flag and, when invalid, a
`validationError` describing why.

```
hostinger wordpress installations list [flags]
```

### Options

```
      --domain string      Filter by domain name (exact match)
  -h, --help               help for list
      --ownership string   Filter by ownership type. Defaults to "owned". Use "all" to include both owned and managed installations. (one of: owned, managed, all)
      --username string    Filter by specific username
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

