## hostinger agency-hosting databases create-plan-website

Create Agency Plan website database

### Synopsis

Creates a MySQL database with a dedicated user for an Agency Plan website.

The database name, username, and password must all be provided by the caller.

```
hostinger agency-hosting databases create-plan-website <website_uid> [flags]
```

### Options

```
      --database-name string   Database name to create (alphanumeric characters).
      --database-user string   Database username to create alongside the database (alphanumeric characters).
  -h, --help                   help for create-plan-website
      --password string        Password for the database user (requires mixed case, letters, and numbers).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

