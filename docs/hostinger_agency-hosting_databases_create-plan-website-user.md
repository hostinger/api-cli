## hostinger agency-hosting databases create-plan-website-user

Create Agency Plan website database user

### Synopsis

Creates a user for an existing database on an Agency Plan website.

Each database supports a single non-system user; creating a user for a database that already has one fails.

```
hostinger agency-hosting databases create-plan-website-user <website_uid> <database_name> [flags]
```

### Options

```
      --database-user string   Database username to create (alphanumeric and underscores).
  -h, --help                   help for create-plan-website-user
      --host string            Host the user connects from (IPv4, IPv6, % wildcard, or localhost). Defaults to localhost.
      --password string        Password for the database user (requires mixed case, letters, and numbers).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

