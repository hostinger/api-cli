## hostinger hosting databases change-password

Change database password

### Synopsis

Changes the password for the specified database user.

The database name must be the full name returned by the list databases endpoint.
The password must also be updated in any website configuration that uses this database.

```
hostinger hosting databases change-password <username> <name> [flags]
```

### Options

```
  -h, --help              help for change-password
      --password string   New database user password.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

