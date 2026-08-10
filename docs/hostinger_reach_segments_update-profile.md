## hostinger reach segments update-profile

Update a profile segment

### Synopsis

Rename a segment and/or replace the conditions that define it.

`name` is always required. Omit `conditions` to rename without touching the conditions;
supply them and they replace the existing set entirely rather than being merged into it.
Contacts are never modified, but which of them match the segment can change immediately.

```
hostinger reach segments update-profile <profile-uuid> <segment-uuid> [flags]
```

### Options

```
      --conditions string   Replaces the existing conditions entirely. Omit to keep the current ones. (JSON)
  -h, --help                help for update-profile
      --logic string        How to combine multiple conditions. Required when conditions are given. (one of: AND, OR)
      --name string         
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

