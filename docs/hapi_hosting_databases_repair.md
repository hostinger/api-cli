## hapi hosting databases repair

Repair database

### Synopsis

Repairs corrupted database tables asynchronously.

Use when database errors, crashes, or corruption are reported.
The database name must be the full name returned by the list databases endpoint.

```
hapi hosting databases repair <username> <name> [flags]
```

### Options

```
  -h, --help   help for repair
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting databases](hapi_hosting_databases.md)	 - Databases commands

