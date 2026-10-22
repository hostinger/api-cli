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

### Examples

```
  hostinger reach segments list-filter-attributes 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-filter-attributes
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

