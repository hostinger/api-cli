## hostinger hosting databases create

Create account database

### Synopsis

Creates a database with a database user and password for the specified account.

The database name and user are automatically prefixed with the account username when needed.

```
hostinger hosting databases create <username> [flags]
```

### Examples

```
  hostinger hosting databases create u123456789 --name u123456789_test_db --password 'Str0ngP@ssword!' --user u123456789_admin --website-domain example.com
```

### Options

```
  -h, --help                    help for create
      --name string             Database name. If the account username prefix is omitted, it is added automatically.
      --password string         Database user password.
      --user string             Database user. If the account username prefix is omitted, it is added automatically.
      --website-domain string   Website domain assigned to the database.
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

