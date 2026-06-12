## hapi reach segments list-profile-contacts

List profile segment contacts

### Synopsis

Retrieve contacts associated with a specific segment for a given profile.

This endpoint allows you to fetch and filter contacts that belong to a particular segment,
identified by its UUID, scoped to a specific profile.

```
hapi reach segments list-profile-contacts <profile-uuid> <segment-uuid> [flags]
```

### Options

```
  -h, --help           help for list-profile-contacts
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

