## hostinger hosting databases delete

Delete account database

### Synopsis

Permanently deletes a database and its remote connections.

The database name must be the full name returned by the list databases endpoint.

```
hostinger hosting databases delete <username> <name> [flags]
```

### Examples

```
  hostinger hosting databases delete u123456789 u123456789_test_db
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

