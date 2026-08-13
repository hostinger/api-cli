## hostinger agency-hosting databases delete-website-user

Delete website database user

### Synopsis

Permanently deletes a database user from an Agency Plan website database, revoking all access it had.

The operation is idempotent: deleting a user that does not exist succeeds without error.

```
hostinger agency-hosting databases delete-website-user <website_uid> <database_name> <database_user_name> [flags]
```

### Options

```
  -h, --help   help for delete-website-user
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

