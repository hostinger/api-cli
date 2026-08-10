## hostinger reach segments list-profile

List profile segments

### Synopsis

Get a paginated list of the segments defined in a profile.

Each entry carries the number of contacts currently matching it, which is recalculated on
read rather than stored. Use `count_type` to count either every matching contact or only
the subscribed ones.

```
hostinger reach segments list-profile <profile-uuid> [flags]
```

### Options

```
      --count-type string   Which matching contacts to count for each segment (one of: all, subscribed) (default "all")
  -h, --help                help for list-profile
      --page int            Page number
      --per-page int        Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

