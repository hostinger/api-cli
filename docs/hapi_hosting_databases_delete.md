## hapi hosting databases delete

Delete account database

### Synopsis

Permanently deletes a database and its remote connections.

The database name must be the full name returned by the list databases endpoint.

```
hapi hosting databases delete <username> <name> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting databases](hapi_hosting_databases.md)	 - Databases commands

