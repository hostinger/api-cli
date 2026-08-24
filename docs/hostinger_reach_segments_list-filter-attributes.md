## hostinger reach segments list-filter-attributes

List segment filter attributes

### Synopsis

List every attribute a segment condition can filter on, with the operators each attribute
accepts, the value format they expect and, where the value is constrained, the allowed
values.

The list is profile specific: it includes the profile's custom contact fields, its tags and
its 20 most recently published campaigns, so the valid attributes cannot be hardcoded. Read
it before creating or updating a segment to discover the valid `attribute`, `operator` and
`value` combinations.

```
hostinger reach segments list-filter-attributes <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for list-filter-attributes
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

