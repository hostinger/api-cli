## hostinger reach segments profile

Get profile segment details

### Synopsis

Get a single segment of a profile, including the conditions that define it.

To retrieve the contacts currently matching those conditions, use the segment contacts
endpoint instead.

```
hostinger reach segments profile <profile-uuid> <segment-uuid> [flags]
```

### Examples

```
  hostinger reach segments profile 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for profile
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

