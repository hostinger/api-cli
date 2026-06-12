## hapi hosting databases list

List account databases

### Synopsis

Returns a paginated list of databases for the specified account.

Use the domain and is_assigned filters to find databases assigned to a specific domain.

```
hapi hosting databases list <username> [flags]
```

### Options

```
      --domain string   Filter by domain name (exact match)
  -h, --help            help for list
      --is-assigned     When used with domain, return only databases assigned to that domain.
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --search string   Search databases by name, user, or creation date.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting databases](hapi_hosting_databases.md)	 - Databases commands

