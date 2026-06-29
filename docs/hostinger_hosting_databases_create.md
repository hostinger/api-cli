## hostinger hosting databases create

Create account database

### Synopsis

Creates a database with a database user and password for the specified account.

The database name and user are automatically prefixed with the account username when needed.

```
hostinger hosting databases create <username> [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

