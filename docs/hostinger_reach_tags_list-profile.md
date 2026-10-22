## hostinger reach tags list-profile

List profile tags

### Synopsis

Get all tags defined in a profile.

Tags are the way contacts are grouped in Reach, and can be used to filter the contact
list or to build segments.

```
hostinger reach tags list-profile <profile-uuid> [flags]
```

### Examples

```
  hostinger reach tags list-profile 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-profile
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

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

