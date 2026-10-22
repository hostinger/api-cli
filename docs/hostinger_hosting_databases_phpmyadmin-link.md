## hostinger hosting databases phpmyadmin-link

Get phpMyAdmin link

### Synopsis

Returns a direct sign-on link to phpMyAdmin for the specified database.

Use this when a visual database interface is needed for SQL queries, imports, exports, or table management.
The database name must be the full name returned by the list databases endpoint.

```
hostinger hosting databases phpmyadmin-link <username> <name> [flags]
```

### Examples

```
  hostinger hosting databases phpmyadmin-link u123456789 u123456789_test_db
```

### Options

```
  -h, --help   help for phpmyadmin-link
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

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

