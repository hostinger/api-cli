## hostinger hosting databases change-password

Change database password

### Synopsis

Changes the password for the specified database user.

The database name must be the full name returned by the list databases endpoint.
The password must also be updated in any website configuration that uses this database.

```
hostinger hosting databases change-password <username> <name> [flags]
```

### Examples

```
  hostinger hosting databases change-password u123456789 u123456789_test_db --password 'Secu4ePa$$wor!D'
```

### Options

```
  -h, --help              help for change-password
      --password string   New database user password.
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

