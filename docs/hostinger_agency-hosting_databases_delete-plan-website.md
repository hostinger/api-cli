## hostinger agency-hosting databases delete-plan-website

Delete Agency Plan website database

### Synopsis

Permanently deletes a MySQL database and all its data from an Agency Plan website, including its users.

The operation is idempotent: deleting a database that does not exist succeeds without error.

```
hostinger agency-hosting databases delete-plan-website <website_uid> <database_name> [flags]
```

### Options

```
  -h, --help   help for delete-plan-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

