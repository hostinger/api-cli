## hapi reach segments list-contacts

List segment contacts

### Synopsis

Retrieve contacts associated with a specific segment.

This endpoint allows you to fetch and filter contacts that belong to a particular segment,
identified by its UUID.

```
hapi reach segments list-contacts <segment-uuid> [flags]
```

### Options

```
  -h, --help           help for list-contacts
      --page int       Page number
      --per-page int   Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi reach segments](hapi_reach_segments.md)	 - Segments commands

